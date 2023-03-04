package render

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var templates = make(map[string]*template.Template)


func init(){

	projectRoot, err := os.Getwd()
	if err != nil {
		log.Panicln("failed to initialise project root", err)
	}

	templatesDir := projectRoot + `\templates`
	templatesMask := templatesDir + `\[^_]*.html`
	templateFiles, err := filepath.Glob(templatesMask)

	if err != nil {
		log.Panicf("failed to glob templates using mask %v. \nError details: %w", templatesMask, err)
	}

	log.Println("templateFiles", templateFiles)

	for _, f := range templateFiles {
		log.Println("caching", f)
		tp, err := template.ParseFiles(f, templatesDir + `\_base.html`)

		if err != nil {
			log.Panicf("error parsing template %v.\nError details: %w", tp, err)
		}
		templates[filepath.Base(f)] = tp
	}

}

func RenderTemplate(w http.ResponseWriter, t string) {
	tp, ok := templates[t]

	if !ok {
		log.Panicln("template", t, "was not found in cache!")
	}

	err := tp.Execute(w, nil)
	if err != nil {
		log.Panicf("error while executing template %v. \nError details: %w", tp, err)
	}

}
