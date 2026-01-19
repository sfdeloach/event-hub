package main

import (
	"net/http"

	"github.com/a-h/templ"

	"server/views"
)

func main() {
	http.Handle("/", templ.Handler(views.Index()))

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.ListenAndServe(":8080", nil)
}
