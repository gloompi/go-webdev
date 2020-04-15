package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type handler int

type data struct {
	Title string
	Date  time.Time
}

var tpl *template.Template
var fm = template.FuncMap{
	"fdateMDY": func(t time.Time) string {
		return t.Format("02-01-2006")
	},
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*"))
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Gloompi-Key", "super secret hashed data")
	w.Header().Set("Content-Type", "text/html; charset=utlf-8")

	switch r.RequestURI {
	case "/":
		indexHandler(w, tpl)
	case "/about":
		aboutHandler(w, tpl)
	default:
		noMatchHandler(w, tpl)
	}
}

func main() {
	var h handler
	http.ListenAndServe(":8080", h)
}
