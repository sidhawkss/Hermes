package main

import (
	"fmt"
	"net/http"
	"html/template"
	"Hermes/pkg/data"
	"Hermes/pkg/conn"
	"Hermes/pkg/tools"
	"Hermes/pkg/events"
	"Hermes/pkg/middleware"
)

func HomeHandle(w http.ResponseWriter, r *http.Request){
	page, err := template.ParseFiles("static/home.html");
	events.TemplateEvent(err);

	page.Execute(w, nil);
}

func ComputerHandle(w http.ResponseWriter, r *http.Request){
	page, err := template.ParseFiles("static/computer.html");
	events.TemplateEvent(err);

	var machines []data.Machine = tools.ReadData();
	page.Execute(w, machines);
}

func ComputerAddHandle(w http.ResponseWriter, r *http.Request){
	page, err := template.ParseFiles("static/add.html");
	events.TemplateEvent(err);

	if r.Method != http.MethodPost {
		page.Execute(w,nil);
	} else {
		r.ParseForm();
		tools.WriteData(r.FormValue("hostname"),r.FormValue("ip"),r.FormValue("country"),r.FormValue("username"),r.FormValue("os"));
		page.Execute(w,1);
	}
}

func InteractHandle(w http.ResponseWriter, r *http.Request){
	page, err := template.ParseFiles("static/interact.html");
	events.TemplateEvent(err);

	var currentId string = r.PathValue("id");
	page.Execute(w, tools.SelectById(currentId));
}

func InteractReceiverHandle(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost{
		r.ParseForm();
		data := conn.Data{CommandString: r.FormValue("data"), ConnectionType: r.PathValue("typeid")};
		response := conn.SendDataTarget(data,r.PathValue("id"));
		
		w.Write([]byte(response));
	}
}


func main(){
	mux := http.NewServeMux()
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))));
	mux.HandleFunc("/", middleware.Authentication(HomeHandle));
	mux.HandleFunc("/computer", middleware.Authentication(ComputerHandle));
	mux.HandleFunc("/computer/add", middleware.Authentication(ComputerAddHandle));
	mux.HandleFunc("/interact/{id}", middleware.Authentication(InteractHandle));
	mux.HandleFunc("/interact/{id}/receiver/{typeid}", middleware.Authentication(InteractReceiverHandle));

	fmt.Println("[MAIN][SUCCESS] Running process in port 3000");
	http.ListenAndServe(":3000",mux);
}
