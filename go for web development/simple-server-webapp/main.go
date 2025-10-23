package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// handler-serving index.html

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error loading the page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// handler - serving helloHandel

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from go backend!")
}

// handler - serving and processing form
func formHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/form.html")
	if err != nil {
		http.Error(w, "error in loading form", http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		message := fmt.Sprintf("Welcome,%s!", name)
		tmpl.Execute(w, message)
		return
	}
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
