package handlers

import (
	"log"
	"net/http"

	"github.com/koleaby4/go-udemy-2/cmd/web/pkg/config"
	"github.com/koleaby4/go-udemy-2/cmd/web/pkg/render"
)

var Repo * Repository

type Repository struct {
	App *config.AppConfig
}

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("serving Home page")
	render.RenderTemplate(w, "home.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	log.Println("serving About page")
	render.RenderTemplate(w, "about.html")
}