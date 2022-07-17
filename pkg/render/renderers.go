package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/pankaj-nikam/go_bookings/pkg/config"
	"github.com/pankaj-nikam/go_bookings/pkg/models"
)

var app *config.AppConfig

//NewTemplates sets the config for templates package.
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, templatePath string, data *models.TemplateData) {
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
		log.Fatal("could get template from template cache")
	}

	buf := new(bytes.Buffer)

	data = AddDefaultData(data)

	err := tmpl.Execute(buf, data)
	if err != nil {
		log.Println(err)
	}

	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//get all the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
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

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
