package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/pankaj-nikam/go_bookings/internal/config"
	"github.com/pankaj-nikam/go_bookings/internal/models"
)

var app *config.AppConfig

var pathToTemplates = "./templates"

// NewTemplates sets the config for templates package.
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, templatePath string, data *models.TemplateData) error {
	var tc map[string]*template.Template

	if app.UseCache {
		//create a template cache.
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//get requested template from cache.
	tmpl, ok := tc[templatePath]
	if !ok {
		log.Println("could get template from template cache")
		return errors.New("cannot get template from cache")
	}

	buf := new(bytes.Buffer)

	data = AddDefaultData(data, r)

	err := tmpl.Execute(buf, data)
	if err != nil {
		log.Println(err)
		return err
	}

	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//get all the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		log.Println("error:", err)
		return myCache, err
	}
	//range through all files which end with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		log.Println("current page:", name)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
