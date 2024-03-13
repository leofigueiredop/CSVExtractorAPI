package main

import (
	"CSVExtractor/handlers"
	"CSVExtractor/services"
	"log"
	"net/http"
)

func main() {

	services.LoadCSVs()

	http.HandleFunc("/search", handlers.SearchHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
