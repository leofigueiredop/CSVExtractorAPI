package main

import (
	"CSVExtractor/handlers"
	"CSVExtractor/services"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Cria um MultiWriter
	mw := io.MultiWriter(os.Stdout, file)

	// Passa o MultiWriter para o log
	log.SetOutput(mw)

	//services.LoadAllCSVs("files")

	services.LoadCSVToDB_CB("files/AILOSDB/BASE_AILOS.csv")
	log.Println("FINALIZOU")

	http.HandleFunc("/search/cadastroBasico", handlers.SearchHandlerPessoa)
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
