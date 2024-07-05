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
	page, err := template.ParseFiles("static/home.html")
	if err != nil {
		fmt.Println("Error: Template parsing.")
	}

	page.Execute(w, nil)
}

func ComputerHandle(w http.ResponseWriter, r *http.Request){
	page, err := template.ParseFiles("static/computer.html")
	if err != nil {
		fmt.Println("Error: Template parsing.")
	}

	var d []data.Machine = operations.ReadData();
	page.Execute(w, d);
}

func ComputerAddHandle(w http.ResponseWriter, r *http.Request){
	page, err := template.ParseFiles("static/add.html")
	if err != nil {
		fmt.Println("Error: Template parsing.")
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
	page, err := template.ParseFiles("static/interact.html")
	if err != nil {
		fmt.Println("Error: Template parsing.")
	}

	page.Execute(w, nil)
}

func InteractReceiverHandle(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost{
		r.ParseForm();
		connectionType := r.PathValue("typeid");
		secStr := enc.Encrypt(r.FormValue("data"));
		resp := conn.SendDataTarget(secStr, connectionType);
		
		w.Write([]byte(resp+" "+secStr));
	}
}

func main(){
	mux := http.NewServeMux()
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))));
	mux.HandleFunc("/", HomeHandle);
	mux.HandleFunc("/computer", ComputerHandle);
	mux.HandleFunc("/computer/add", ComputerAddHandle);
	mux.HandleFunc("/interact", InteractHandle);
	mux.HandleFunc("/interact/receiver/{typeid}", InteractReceiverHandle);

	fmt.Println("Running")
	http.ListenAndServe(":3000", mux)
}

