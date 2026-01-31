package main

import (
	"net/http"
	"server/database"
	"server/handlers"
)

func main() {
	db := database.Init()
	h := &handlers.Handler{DB: db}

	http.Handle("/", http.RedirectHandler("/home", http.StatusFound))
	http.HandleFunc("GET /home", h.Home)
	http.HandleFunc("GET /events", h.Events)
	http.HandleFunc("GET /events/create", h.CreateEvent)
	http.HandleFunc("POST /events/create", h.PostCreateEvent)
	http.HandleFunc("GET /events/edit/{id}", h.EditEvent)
	http.HandleFunc("PUT /events/edit/{id}", h.PutEditEvent)
	http.HandleFunc("DELETE /events/delete/{id}", h.DeleteEvent)
	http.HandleFunc("GET /events/categories", h.EventCategories)
	http.HandleFunc("GET /events/categories/create", h.CreateEventCategory)
	http.HandleFunc("POST /events/categories/create", h.PostCreateEventCategory)
	http.HandleFunc("GET /events/categories/edit/{id}", h.EditEventCategory)
	http.HandleFunc("PUT /events/categories/edit/{id}", h.PutEditEventCategory)
	http.HandleFunc("DELETE /events/categories/delete/{id}", h.DeleteEventCategory)
	http.HandleFunc("GET /login", h.Login)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.ListenAndServe(":8080", nil)
}
