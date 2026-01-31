package handlers

import (
	"net/http"
	"server/models"
	"server/views"
	"strconv"
	"strings"
	"time"

	"github.com/a-h/templ"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func isHtmxRequest(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}

// render handles rendering HTMX partial or full page based on request type
func render(w http.ResponseWriter, r *http.Request, content, full templ.Component) {
	var err error
	if isHtmxRequest(r) {
		err = content.Render(r.Context(), w)
	} else {
		err = full.Render(r.Context(), w)
	}
	if err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}

// redirectToCategoryList renders category list for HTMX or redirects for regular requests
func (h *Handler) redirectToCategoryList(w http.ResponseWriter, r *http.Request) {
	if isHtmxRequest(r) {
		var categories []models.EventCategory
		if err := h.DB.Find(&categories).Error; err != nil {
			http.Error(w, "Failed to load categories", http.StatusInternalServerError)
			return
		}
		if err := views.EventCategoriesContent(categories).Render(r.Context(), w); err != nil {
			http.Error(w, "Failed to render page", http.StatusInternalServerError)
		}
	} else {
		http.Redirect(w, r, "/events/categories", http.StatusSeeOther)
	}
}

// redirectToEventList renders event list for HTMX or redirects for regular requests
func (h *Handler) redirectToEventList(w http.ResponseWriter, r *http.Request) {
	if isHtmxRequest(r) {
		now := time.Now()

		var visibleEvents []models.Event
		if err := h.DB.Where("always_visible = ? OR (on_air <= ? AND off_air >= ?)", true, now, now).Find(&visibleEvents).Error; err != nil {
			http.Error(w, "Failed to load events", http.StatusInternalServerError)
			return
		}

		var offAirEvents []models.Event
		if err := h.DB.Where("always_visible = ? AND (on_air > ? OR off_air < ?)", false, now, now).Find(&offAirEvents).Error; err != nil {
			http.Error(w, "Failed to load off-air events", http.StatusInternalServerError)
			return
		}

		if err := views.EventsContent(visibleEvents, offAirEvents).Render(r.Context(), w); err != nil {
			http.Error(w, "Failed to render page", http.StatusInternalServerError)
		}
	} else {
		http.Redirect(w, r, "/events", http.StatusSeeOther)
	}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	render(w, r, views.HomeContent(), views.Home())
}

func (h *Handler) Events(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	// Query visible events: always visible OR within the on_air/off_air window
	var visibleEvents []models.Event
	if err := h.DB.Where("always_visible = ? OR (on_air <= ? AND off_air >= ?)", true, now, now).Find(&visibleEvents).Error; err != nil {
		http.Error(w, "Failed to load events", http.StatusInternalServerError)
		return
	}

	// Query off-air events: NOT always visible AND outside the on_air/off_air window
	var offAirEvents []models.Event
	if err := h.DB.Where("always_visible = ? AND (on_air > ? OR off_air < ?)", false, now, now).Find(&offAirEvents).Error; err != nil {
		http.Error(w, "Failed to load off-air events", http.StatusInternalServerError)
		return
	}

	render(w, r, views.EventsContent(visibleEvents, offAirEvents), views.Events(visibleEvents, offAirEvents))
}

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var categories []models.EventCategory
	if err := h.DB.Find(&categories).Error; err != nil {
		http.Error(w, "Failed to load categories", http.StatusInternalServerError)
		return
	}
	render(w, r, views.CreateEventContent(categories), views.CreateEvent(categories))
}

func (h *Handler) PostCreateEvent(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	// Parse required fields
	title := strings.TrimSpace(r.FormValue("title"))
	if title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	when := strings.TrimSpace(r.FormValue("when"))
	if when == "" {
		http.Error(w, "When is required", http.StatusBadRequest)
		return
	}

	where := strings.TrimSpace(r.FormValue("where"))
	if where == "" {
		http.Error(w, "Where is required", http.StatusBadRequest)
		return
	}

	// Parse category ID
	categoryIDStr := r.FormValue("category")
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid category", http.StatusBadRequest)
		return
	}

	// Check if always visible
	alwaysVisible := r.FormValue("always_visible") == "on"

	// Parse on_air and off_air datetimes (only required if not always visible)
	var onAir, offAir time.Time
	if !alwaysVisible {
		onAir, err = time.ParseInLocation("2006-01-02T15:04", r.FormValue("on_air_at"), time.Local)
		if err != nil {
			http.Error(w, "Invalid on air date/time", http.StatusBadRequest)
			return
		}

		offAir, err = time.ParseInLocation("2006-01-02T15:04", r.FormValue("off_air_at"), time.Local)
		if err != nil {
			http.Error(w, "Invalid off air date/time", http.StatusBadRequest)
			return
		}
	}

	event := models.Event{
		Title:           title,
		Description:     strings.TrimSpace(r.FormValue("description")),
		AlwaysVisible:   alwaysVisible,
		When:            when,
		Where:           where,
		OnAir:           onAir,
		OffAir:          offAir,
		EventCategoryID: uint(categoryID),
	}

	if err := h.DB.Create(&event).Error; err != nil {
		http.Error(w, "Failed to create event", http.StatusInternalServerError)
		return
	}

	h.redirectToEventList(w, r)
}

func (h *Handler) EditEvent(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var event models.Event
	if err := h.DB.First(&event, id).Error; err != nil {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}

	var categories []models.EventCategory
	if err := h.DB.Find(&categories).Error; err != nil {
		http.Error(w, "Failed to load categories", http.StatusInternalServerError)
		return
	}

	render(w, r, views.EditEventContent(event, categories), views.EditEvent(event, categories))
}

func (h *Handler) PutEditEvent(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var event models.Event
	if err := h.DB.First(&event, id).Error; err != nil {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	// Parse required fields
	title := strings.TrimSpace(r.FormValue("title"))
	if title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	when := strings.TrimSpace(r.FormValue("when"))
	if when == "" {
		http.Error(w, "When is required", http.StatusBadRequest)
		return
	}

	where := strings.TrimSpace(r.FormValue("where"))
	if where == "" {
		http.Error(w, "Where is required", http.StatusBadRequest)
		return
	}

	// Parse category ID
	categoryIDStr := r.FormValue("category")
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid category", http.StatusBadRequest)
		return
	}

	// Check if always visible
	alwaysVisible := r.FormValue("always_visible") == "on"

	// Parse on_air and off_air datetimes (only required if not always visible)
	var onAir, offAir time.Time
	if !alwaysVisible {
		onAir, err = time.ParseInLocation("2006-01-02T15:04", r.FormValue("on_air_at"), time.Local)
		if err != nil {
			http.Error(w, "Invalid on air date/time", http.StatusBadRequest)
			return
		}

		offAir, err = time.ParseInLocation("2006-01-02T15:04", r.FormValue("off_air_at"), time.Local)
		if err != nil {
			http.Error(w, "Invalid off air date/time", http.StatusBadRequest)
			return
		}
	}

	// Update event fields
	event.Title = title
	event.Description = strings.TrimSpace(r.FormValue("description"))
	event.When = when
	event.Where = where
	event.AlwaysVisible = alwaysVisible
	event.OnAir = onAir
	event.OffAir = offAir
	event.EventCategoryID = uint(categoryID)

	if err := h.DB.Save(&event).Error; err != nil {
		http.Error(w, "Failed to update event", http.StatusInternalServerError)
		return
	}

	h.redirectToEventList(w, r)
}

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var event models.Event
	if err := h.DB.First(&event, id).Error; err != nil {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}

	if err := h.DB.Delete(&event).Error; err != nil {
		http.Error(w, "Failed to delete event", http.StatusInternalServerError)
		return
	}

	// Return empty response - HTMX will replace the element with nothing
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) EventCategories(w http.ResponseWriter, r *http.Request) {
	var categories []models.EventCategory
	if err := h.DB.Find(&categories).Error; err != nil {
		http.Error(w, "Failed to load categories", http.StatusInternalServerError)
		return
	}
	render(w, r, views.EventCategoriesContent(categories), views.EventCategories(categories))
}

func (h *Handler) CreateEventCategory(w http.ResponseWriter, r *http.Request) {
	render(w, r, views.CreateEventCategoryContent(), views.CreateEventCategory())
}

func (h *Handler) PostCreateEventCategory(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	category := strings.TrimSpace(r.FormValue("category"))
	if category == "" {
		http.Error(w, "Category is required", http.StatusBadRequest)
		return
	}

	if err := h.DB.Create(&models.EventCategory{Category: category}).Error; err != nil {
		http.Error(w, "Failed to create category", http.StatusInternalServerError)
		return
	}

	h.redirectToCategoryList(w, r)
}

func (h *Handler) EditEventCategory(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var category models.EventCategory
	if err := h.DB.First(&category, id).Error; err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}
	render(w, r, views.EditEventCategoryContent(category), views.EditEventCategory(category))
}

func (h *Handler) PutEditEventCategory(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var category models.EventCategory
	if err := h.DB.First(&category, id).Error; err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	newCategory := strings.TrimSpace(r.FormValue("category"))
	if newCategory == "" {
		http.Error(w, "Category is required", http.StatusBadRequest)
		return
	}

	category.Category = newCategory
	if err := h.DB.Save(&category).Error; err != nil {
		http.Error(w, "Failed to update category", http.StatusInternalServerError)
		return
	}

	h.redirectToCategoryList(w, r)
}

func (h *Handler) DeleteEventCategory(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var category models.EventCategory
	if err := h.DB.First(&category, id).Error; err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	if err := h.DB.Delete(&category).Error; err != nil {
		http.Error(w, "Failed to delete category", http.StatusInternalServerError)
		return
	}

	// Return empty response - HTMX will replace the element with nothing
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	render(w, r, views.LoginContent(), views.Login())
}
