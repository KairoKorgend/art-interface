package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/art-interface/assets/", http.StripPrefix("/art-interface/assets", fs))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/decoder", processFormHandler)
	fmt.Println("Server starting on port 8080...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func processFormHandler(w http.ResponseWriter, r *http.Request) {

	encodeModeValue := r.FormValue("typeValue")
	encodeModeBool := encodeModeValue == "encode"
	input := r.FormValue("textValue")

	result, err := ProcessLine(input, encodeModeBool)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := struct {
		Input  string
		Result string
	}{
		Input:  input,
		Result: result,
	}

	w.WriteHeader(http.StatusAccepted)

	tmpl.ExecuteTemplate(w, "index.html", data)
}
