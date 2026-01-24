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

func NewEvent(w http.ResponseWriter, r *http.Request) {
	if isHtmxRequest(r) {
		views.NewEventContent().Render(r.Context(), w)
	} else {
		views.NewEvent().Render(r.Context(), w)
	}
}