package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mo0Oonnn/bookings/internal/config"
	"github.com/mo0Oonnn/bookings/internal/handlers"
	"github.com/mo0Oonnn/bookings/internal/models"
	"github.com/mo0Oonnn/bookings/internal/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application module
func main() {
	gob.Register(models.Reservation{})

	app.InProduction = false

	session = scs.New()
	session.Lifetime = time.Hour * 24
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc := render.CreateTemplateCache()

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.SetRepo(repo)

	render.SetConfig(&app)

	srv := http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	fmt.Printf("Starting application on port %s\n", portNumber)

	err := srv.ListenAndServe()
	log.Fatal(err)
}
