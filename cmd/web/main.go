package main

import (
	"log"
	"net/http"

	"github.com/koleaby4/go-udemy-2/cmd/web/pkg/handlers"
)

const portNumber = ":8080"


func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Printf("starting the application on port %v", portNumber)
	http.ListenAndServe(portNumber, nil)
}
