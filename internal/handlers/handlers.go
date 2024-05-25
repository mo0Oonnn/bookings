package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mo0Oonnn/bookings/internal/config"
	"github.com/mo0Oonnn/bookings/internal/forms"
	"github.com/mo0Oonnn/bookings/internal/models"
	"github.com/mo0Oonnn/bookings/internal/render"

	"github.com/CloudyKit/jet"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func SetRepo(r *Repository) {
	Repo = r
}

func (rep *Repository) Home(w http.ResponseWriter, r *http.Request) {
	data := make(jet.VarMap)

	data.Set("greeting", "hello!")

	render.RenderTemplate(w, r, "home.jet", data)
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "about.jet", jet.VarMap{})
}

// Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(jet.VarMap)
	data.Set("reservation", emptyReservation)
	data.Set("form", forms.NewForm(nil))

	render.RenderTemplate(w, r, "make-reservation.jet", data)
}

// PostReservation handles the posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		DateFrom:  r.Form.Get("date_from"),
		DateTo:    r.Form.Get("date_to"),
		Phone:     r.Form.Get("phone"),
		Email:     r.Form.Get("email"),
		Room:      r.Form.Get("room"),
	}

	form := forms.NewForm(r.PostForm)

	form.IsRequired("first_name", "last_name", "date_from", "date_to", "email")
	form.CheckLength("first_name", 3, r)
	form.CheckLength("last_name", 3, r)
	form.IsEmail("email")

	if !form.IsValid() {
		data := make(jet.VarMap)
		data.Set("reservation", reservation)
		data.Set("form", form)

		render.RenderTemplate(w, r, "make-reservation.jet", data)
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

// SingleRoom renders the room page
func (m *Repository) SingleRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "single-room.jet", jet.VarMap{})
}

// TwoBedRoom renders the room page
func (m *Repository) TwoBedRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "two-bed-room.jet", jet.VarMap{})
}

// DoubleBedRoom renders the room page
func (m *Repository) DoubleBedRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "double-bed-room.jet", jet.VarMap{})
}

// FamilyRoom renders the room page
func (m *Repository) FamilyRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "family-room.jet", jet.VarMap{})
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.jet", jet.VarMap{})
}

// PostAvailability renders the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("Дата въезда: %s, дата выезда: %s", start, end)))
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.jet", jet.VarMap{})
}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("cannot get item from session")
		m.App.Session.Put(r.Context(), "error", "Невозможно получить сводку бронирования")
		m.App.Session.Put(r.Context(), "isThereError", true)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")

	data := make(jet.VarMap)
	data.Set("reservation", reservation)

	render.RenderTemplate(w, r, "reservation-summary.jet", data)
}
