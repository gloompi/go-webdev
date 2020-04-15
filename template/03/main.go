package main

import (
	"html/template"
	"log"
	"os"
	"strings"
	"time"
)

type person struct {
	Name     string
	Lastname string
	Age      int
}

func (p person) GetFullName() string {
	return p.Name + " " + p.Lastname
}

type transport struct {
	Doors  int
	Wheels int
	Model  string
}

type data struct {
	Users      []person
	Transports []transport
	Date       time.Time
}

var tpl *template.Template
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": func(s string) string {
		s = strings.TrimSpace(s)
		s = s[:3]
		return s
	},
	"fdateMDY": func(t time.Time) string {
		return t.Format("02-01-2006")
	},
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("template/*"))
}

func main() {
	Users := []person{
		{Name: "Kuba", Lastname: "Esenzhanov", Age: 24},
		{Name: "Tina", Lastname: "Nurdin kyzy", Age: 23},
		{Name: "Meroj", Lastname: "Uzbek", Age: 22},
		{Name: "Xixi", Lastname: "Vietnamese", Age: 24},
	}

	Transports := []transport{
		{Model: "BMW", Doors: 2, Wheels: 4},
		{Model: "Mers", Doors: 4, Wheels: 4},
		{Model: "Suzuki", Doors: 0, Wheels: 2},
	}

	date := time.Now()

	data := data{Users, Transports, date}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
