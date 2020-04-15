package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func noMatchHandler(w http.ResponseWriter, tpl *template.Template) {
	err := tpl.ExecuteTemplate(w, "404.gohtml", data{Title: "No Matchings", Date: time.Now()})
	if err != nil {
		log.Fatalln(err)
	}
}
