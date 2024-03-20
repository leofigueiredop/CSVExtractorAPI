package main

import (
	"CSVExtractor/services"
	"io"
	"log"
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

	// Carregue todos os arquivos CSV relacionados na memória
	services.LoadAllCSVs("files")

	// Carregue o arquivo "Cadastro Básico" CSV na memória e gere o arquivo JSON
	outputJSONFilePath := "path-to-your-output-directory/output.json"
	services.LoadCSVToMemory_CB("files/AILOSDB/BASE_AILOS.csv", outputJSONFilePath)

	log.Println("FINALIZOU")

	//log.Fatal(http.ListenAndServe(":8080", nil))
}
