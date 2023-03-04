package main

import (
	"log"
	"net/http"

	"github.com/koleaby4/go-udemy-2/cmd/web/pkg/handlers"
	"github.com/koleaby4/go-udemy-2/cmd/web/pkg/render"
)

const portNumber = ":8080"


func main() {

	err := render.InitCache()

	if err != nil {
		log.Panicln(err)
	}


	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Printf("starting the application on port %v", portNumber)
	http.ListenAndServe(portNumber, nil)
}
