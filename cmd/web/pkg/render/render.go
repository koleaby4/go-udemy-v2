package render

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var cache map[string]*template.Template

func InitCache() error{
	templates, err := CachePageTemplates()
	if err != nil {
		return err
	}
	cache = templates
	return nil
}

func CachePageTemplates() (map[string]*template.Template, error) {
	var templates = make(map[string]*template.Template)

	projectRoot, err := os.Getwd()
	if err != nil {
		return templates, err
	}

	templatesDir := projectRoot + `\templates`
	templatesMask := templatesDir + `\[^_]*.html`
	templateFiles, err := filepath.Glob(templatesMask)

	log.Println("templateFiles", templateFiles)

	if err != nil {
		return templates, err
	}

	baseFiles, err := filepath.Glob(templatesDir + `\_base.html`)
	if err != nil {
		return templates, err
	}

	if len(baseFiles) == 0 {
		return templates, errors.New("no base layout files found")
	}

	base := baseFiles[0]

	for _, f := range templateFiles {
		log.Println("caching", f)
		tp, err := template.ParseFiles(f, base)
		if err != nil {
			return templates, err
		}
		templates[filepath.Base(f)] = tp
	}

	return templates, nil
}

func RenderTemplate(w http.ResponseWriter, t string) {
	tp, ok := cache[t]

	if !ok {
		log.Panicln("template", t, "was not found in cache!")
	}

	err := tp.Execute(w, nil)
	if err != nil {
		log.Panicln(err)
	}

}
