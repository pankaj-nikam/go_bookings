package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/pankaj-nikam/go_bookings/internal/config"
	"github.com/pankaj-nikam/go_bookings/internal/handlers"
	"github.com/pankaj-nikam/go_bookings/internal/models"
	"github.com/pankaj-nikam/go_bookings/internal/render"
)

const portNumber = ":3000"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	//What am I going to store in session
	gob.Register(models.Reservation{})

	//Change this to true when in production
	app.IsProduction = false

	session = scs.New()
	session.Lifetime = time.Hour * 24
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.IsProduction
	session.Cookie.HttpOnly = true

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc

	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println("Listening on port", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
