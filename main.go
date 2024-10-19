package main

import (
	"fmt"
	"net/http"
	"html/template"
	"Hermes/pkg/data"
	"Hermes/pkg/conn"
	"Hermes/pkg/enc"
	"Hermes/pkg/operations"
)

func HomeHandle(w http.ResponseWriter, r *http.Request){
	page, err := template.ParseFiles("static/home.html");
	if err != nil {
		fmt.Println("Error: Template parsing.");
	}

	page.Execute(w, nil);
}

func ComputerHandle(w http.ResponseWriter, r *http.Request){
	page, err := template.ParseFiles("static/computer.html");
	if err != nil {
		fmt.Println("Error: Template parsing.");
	}

	var machines []data.Machine = operations.ReadData();
	page.Execute(w, machines);
}

func ComputerAddHandle(w http.ResponseWriter, r *http.Request){
	page, err := template.ParseFiles("static/add.html");
	if err != nil {
		fmt.Println("Error: Template parsing.");
	}

	if r.Method != http.MethodPost {
		page.Execute(w,nil);
	}else{
		r.ParseForm();
		operations.WriteData(r.FormValue("hostname"),r.FormValue("ip"),r.FormValue("country"),r.FormValue("username"),r.FormValue("os"));
		page.Execute(w,1);
	}
	
}

func InteractHandle(w http.ResponseWriter, r *http.Request){
	page, err := template.ParseFiles("static/interact.html");
	if err != nil {
		fmt.Println("Error: Template parsing.");
	}
	
	var currentId string = r.PathValue("id");
	var machines []data.Machine = operations.ReadData();
	var currentMachine data.Machine;
	for i := 0; i < len(machines); i++ {
		if currentId == machines[i].Id {
			currentMachine = machines[i];
		}
	}

	fmt.Println(currentMachine);
	page.Execute(w, currentMachine);
}

func InteractReceiverHandle(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost{
		r.ParseForm();
		data := conn.Data{SecureString: enc.Encrypt(r.FormValue("data")), ConnectionType: r.PathValue("typeid")};
		response := conn.SendDataTarget(data,r.PathValue("id"));
		
		w.Write([]byte(response));
	}
}

func main(){
	mux := http.NewServeMux()
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))));
	mux.HandleFunc("/", HomeHandle);
	mux.HandleFunc("/computer", ComputerHandle);
	mux.HandleFunc("/computer/add", ComputerAddHandle);
	mux.HandleFunc("/interact/{id}", InteractHandle);
	mux.HandleFunc("/interact/{id}/receiver/{typeid}", InteractReceiverHandle);

	fmt.Println("[MAIN][SUCCESS] running process in port 3000")
	http.ListenAndServe(":3000",mux)
}
