package main

import (
	"fmt"
	"net/http"
	"html/template"
	"HermesC2/pkg/data"
	"HermesC2/pkg/operations"
)


func Home(w http.ResponseWriter, r *http.Request){
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
	page.Execute(w, nil);
}


func main(){
	mux := http.NewServeMux()
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/computer", ComputerHandle)
	mux.HandleFunc("/computer/add", ComputerAddHandle)
	fmt.Println("Running")

	http.ListenAndServe(":3000", mux)
}


