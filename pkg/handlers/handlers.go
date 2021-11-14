package handlers

import (
	"net/http"
	"rarebnb/pkg/config"
	"rarebnb/pkg/models"
	"rarebnb/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

//About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again."

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

//Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}

//Reservation renders the make a reservation page and displays form
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}

//Availability renders the search availability room
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
}

//Timothys renders the page for the "Timothy's Chalet"
func (m *Repository) Timothys(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "timothys.page.tmpl", &models.TemplateData{})
}

//HurtYurt renders the page for "The Hurt Yurt"
func (m *Repository) HurtYurt(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "hurtyurt.page.tmpl", &models.TemplateData{})
}
