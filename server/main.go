package main

import (
	"net/http"

	"server/views"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/events", handleEvents)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.ListenAndServe(":8080", nil)
}

func isHtmxRequest(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if isHtmxRequest(r) {
		views.IndexContent().Render(r.Context(), w)
	} else {
		views.Index().Render(r.Context(), w)
	}
}

func handleEvents(w http.ResponseWriter, r *http.Request) {
	if isHtmxRequest(r) {
		views.EventsContent().Render(r.Context(), w)
	} else {
		views.Events().Render(r.Context(), w)
	}
}
