package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func aboutHandler(w http.ResponseWriter, tpl *template.Template) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", data{Title: "About", Date: time.Now()})
	if err != nil {
		log.Fatalln(err)
	}
}
