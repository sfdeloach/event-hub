package handlers

import (
	"net/http"
	"server/views"
)

func isHtmxRequest(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}

func Index(w http.ResponseWriter, r *http.Request) {
	if isHtmxRequest(r) {
		views.IndexContent().Render(r.Context(), w)
	} else {
		views.Index().Render(r.Context(), w)
	}
}

func Events(w http.ResponseWriter, r *http.Request) {
	if isHtmxRequest(r) {
		views.EventsContent().Render(r.Context(), w)
	} else {
		views.Events().Render(r.Context(), w)
	}
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	if isHtmxRequest(r) {
		views.CreateEventContent().Render(r.Context(), w)
	} else {
		views.CreateEvent().Render(r.Context(), w)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if isHtmxRequest(r) {
		views.LoginContent().Render(r.Context(), w)
	} else {
		views.Login().Render(r.Context(), w)
	}
}
