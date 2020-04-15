package main

import (
	"fmt"
	"net/http"
)

type handler int

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var h handler
	http.ListenAndServe(":8080", h)
}
