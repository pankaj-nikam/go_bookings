package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/pankaj-nikam/hello_web/pkg/config"
	"github.com/pankaj-nikam/hello_web/pkg/handlers"
	"github.com/pankaj-nikam/hello_web/pkg/render"
)

const portNumber = ":3000"

var app config.AppConfig
var session *scs.SessionManager

func main() {

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

	app.UseCache = true

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
