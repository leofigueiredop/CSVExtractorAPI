package main

import (
	"CSVExtractor/handlers"
	"CSVExtractor/services"
	"log"
	"net/http"
)

func main() {

	services.LoadAllCSVs("files")

	http.HandleFunc("/search/pep", handlers.SearchHandlerPEP)
	http.HandleFunc("/search/ceis", handlers.SearchHandlerCEIS)
	http.HandleFunc("/search/cnep", handlers.SearchHandlerCNEP)
	http.HandleFunc("/search/AutosInfracaoIbama", handlers.SearchHandlerAutosInfracaoIbama)
	http.HandleFunc("/search/AutosInfracaoICMBIO", handlers.SearchHandlerAutosInfracaoICMBIO)
	http.HandleFunc("/search/TrabalhoEscravo", handlers.SearchHandlerTrabalhoEscravo)
	http.HandleFunc("/search/Suspensaobama", handlers.SearchHandlerSuspensaobama)
	http.HandleFunc("/search/ApreensaoIbama", handlers.SearchHandlerApreensaoIbama)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
