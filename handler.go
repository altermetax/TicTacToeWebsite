package main

import (
	"net/http"
	"html/template"
)

var templates = template.Must(template.ParseFiles("home.html"))

func renderTemplate(w http.ResponseWriter, t string) {
	err := templates.ExecuteTemplate(w, t+".html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func makeStaticPageHandler(page string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, page)
	}
}