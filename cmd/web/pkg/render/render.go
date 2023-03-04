package render

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/koleaby4/go-udemy-2/cmd/web/pkg/models"
)

var cache = make(map[string]*template.Template)


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
		cache[filepath.Base(f)] = tp
	}

}



func AddDefaultData(data *models.TemplateData) *models.TemplateData{
	return data
}


func RenderTemplate(w http.ResponseWriter, t string, data *models.TemplateData) {
	tp, ok := cache[t]

	if !ok {
		log.Panicln("template", t, "was not found in cache!")
	}

	data = AddDefaultData(data)

	err := tp.Execute(w, data)

	if err != nil {
		log.Panicf("error while executing template %v. \nError details: %w", tp, err)
	}

}
