package handlers

import (
	"CSVExtractor/services"
	"encoding/json"
	"net/http"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get("term")
	layout := r.URL.Query().Get("layout")

	// execute a search function here
	result := services.Search(term, layout)

	// transform result into json
	jsonResponse, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set the header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
