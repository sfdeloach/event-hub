package handlers

import (
	"net/http"
	"server/models"
	"server/views"

	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func isHtmxRequest(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if isHtmxRequest(r) {
		views.HomeContent().Render(r.Context(), w)
	} else {
		views.Home().Render(r.Context(), w)
	}
}

func (h *Handler) Events(w http.ResponseWriter, r *http.Request) {
	if isHtmxRequest(r) {
		views.EventsContent().Render(r.Context(), w)
	} else {
		views.Events().Render(r.Context(), w)
	}
}

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var categories []models.EventCategory
	h.DB.Find(&categories)

	if isHtmxRequest(r) {
		views.CreateEventContent(categories).Render(r.Context(), w)
	} else {
		views.CreateEvent(categories).Render(r.Context(), w)
	}
}

func (h *Handler) EventCategories(w http.ResponseWriter, r *http.Request) {
	var categories []models.EventCategory
	h.DB.Find(&categories)

	if isHtmxRequest(r) {
		views.EventCategoriesContent(categories).Render(r.Context(), w)
	} else {
		views.EventCategories(categories).Render(r.Context(), w)
	}
}

func (h *Handler) CreateEventCategory(w http.ResponseWriter, r *http.Request) {
	if isHtmxRequest(r) {
		views.CreateEventCategoryContent().Render(r.Context(), w)
	} else {
		views.CreateEventCategory().Render(r.Context(), w)
	}
}

func (h *Handler) PostCreateEventCategory(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	category := r.FormValue("category")
	if category == "" {
		http.Error(w, "Category is required", http.StatusBadRequest)
		return
	}

	result := h.DB.Create(&models.EventCategory{Category: category})
	if result.Error != nil {
		http.Error(w, "Failed to create category", http.StatusInternalServerError)
		return
	}

	if isHtmxRequest(r) {
		var categories []models.EventCategory
		h.DB.Find(&categories)
		views.EventCategoriesContent(categories).Render(r.Context(), w)
	} else {
		http.Redirect(w, r, "/events/categories", http.StatusSeeOther)
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if isHtmxRequest(r) {
		views.LoginContent().Render(r.Context(), w)
	} else {
		views.Login().Render(r.Context(), w)
	}
}
