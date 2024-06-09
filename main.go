package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	http.HandleFunc("/home",
		func(w http.ResponseWriter, r *http.Request) {
			tpl, _ = template.ParseFiles("templates/home.html")
			tpl.Execute(w, nil)
		})

	http.ListenAndServe(":8000", http.DefaultServeMux)
}
 