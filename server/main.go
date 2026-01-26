package main

import (
	"net/http"
	"server/handlers"
)

func main() {
	http.Handle("/", http.RedirectHandler("/home", http.StatusFound))
	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/events", handlers.Events)
	http.HandleFunc("/events/create", handlers.CreateEvent)
	http.HandleFunc("/events/categories", handlers.EventCategories)
	http.HandleFunc("/events/categories/create", handlers.CreateEventCategory)
	http.HandleFunc("/login", handlers.Login)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.ListenAndServe(":8080", nil)
}
