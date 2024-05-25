package main

import (
	"fmt"
	"net/http"
	"html/template"
	"HermesC2/pkg/data"
	"HermesC2/pkg/operations"
)


func Home(w http.ResponseWriter, r *http.Request){
	page, err := template.ParseFiles("static/index.html")
	if err != nil {
		fmt.Println("Error: Template parsing.")
	}

	page.Execute(w, nil)
}

func ComputersHandle(w http.ResponseWriter, r *http.Request){
	page, err := template.ParseFiles("static/computers.html")
	if err != nil {
		fmt.Println("Error: Template parsing.")
	}

	var c []data.Machine = operations.ReadData();
	page.Execute(w, c);
}

func main(){
	mux := http.NewServeMux()
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/computers", ComputersHandle)
	fmt.Println("Running")

	http.ListenAndServe(":3000", mux)
}


