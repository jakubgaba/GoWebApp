package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsin template:", err)
		return
	}
}

var templateCache = make(map[string]*template.Template)

func RenderTemplateTest(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// Check to see if we already have the template in our cache
	_, inMap := templateCache[t]
	if !inMap {
		//need to create the template

		log.Println("Creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		//we have the temp
		log.Println("using cached template")
	}

	tmpl = templateCache[t]

	err = tmpl.Execute(w, nil)

	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	//parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	//add template to cache(map)

	templateCache[t] = tmpl

	return nil
}
