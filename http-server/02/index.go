package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func indexHandler(w http.ResponseWriter, tpl *template.Template) {
	d := data{Title: "Home", Date: time.Now()}
	err := tpl.ExecuteTemplate(w, "index.gohtml", d)
	if err != nil {
		log.Fatalln(err)
	}
}
