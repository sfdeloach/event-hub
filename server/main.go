package main

import (
	"net/http"
	"server/handlers"
)

func main() {
	http.Handle("/", http.RedirectHandler("/home", http.StatusFound))
	http.HandleFunc("/home", handlers.Index)
	http.HandleFunc("/events", handlers.Events)
	http.HandleFunc("/events/new", handlers.NewEvent)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.ListenAndServe(":8080", nil)
}
