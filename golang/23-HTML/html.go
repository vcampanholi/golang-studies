package main

import (
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Teste HTTP"))
}

func main() {
	http.HandleFunc("/home", home)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
