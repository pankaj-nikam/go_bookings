package handlers

import (
	"net/http"

	"github.com/pankaj-nikam/go_bookings/pkg/config"
	"github.com/pankaj-nikam/go_bookings/pkg/models"
	"github.com/pankaj-nikam/go_bookings/pkg/render"
)

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//Repo the repository used by handlers
var Repo *Repository

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//Perform some business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again!"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	//Send the data received to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

//Generals renders the generals room page.
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "generals.page.tmpl", &models.TemplateData{})
}

//Majors renders the majors room page.
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors.page.tmpl", &models.TemplateData{})
}

//SearchAvailability renders the make a reservation page and displays form.
func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
}

//Reservation renders the contact.
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}

//Reservation renders the make a reservation page and displays form.
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}
