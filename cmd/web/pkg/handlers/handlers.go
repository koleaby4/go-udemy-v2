package handlers

import (
	"log"
	"net/http"

	"github.com/koleaby4/go-udemy-2/cmd/web/pkg/models"
	"github.com/koleaby4/go-udemy-2/cmd/web/pkg/render"
)


func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("serving Home page")
	render.RenderTemplate(w, "home.html", &models.TemplateData{})
}

func About(w http.ResponseWriter, r *http.Request) {
	log.Println("serving About page")

	args := map[string]string {
		"test": "Hello from args",
	}

	render.RenderTemplate(w, "about.html", &models.TemplateData{StringMap : args})
}
