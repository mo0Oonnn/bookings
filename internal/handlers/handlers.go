package handlers

import (
	"fmt"
	"net/http"

	"github.com/mo0Oonnn/bookings/internal/config"
	"github.com/mo0Oonnn/bookings/internal/forms"
	"github.com/mo0Oonnn/bookings/internal/models"
	"github.com/mo0Oonnn/bookings/internal/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates the repository type
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["prompt"] = "Also you can go to the home page"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make_reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
	})
}

// PostReservation handles the posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
}

// SingleRoom renders the room page
func (m *Repository) SingleRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "single_room.page.tmpl", &models.TemplateData{})
}

// TwoBedRoom renders the room page
func (m *Repository) TwoBedRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "two_bed_room.page.tmpl", &models.TemplateData{})
}

// DoubleBedRoom renders the room page
func (m *Repository) DoubleBedRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "double_bed_room.page.tmpl", &models.TemplateData{})
}

// FamilyRoom renders the room page
func (m *Repository) FamilyRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "family_room.page.tmpl", &models.TemplateData{})
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search_availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability renders the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	room := r.Form.Get("room")

	w.Write([]byte(fmt.Sprintf("Дата въезда: %s, дата выезда: %s, выбранный номер: %s", start, end, room)))
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}
