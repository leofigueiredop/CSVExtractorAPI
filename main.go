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
	//services.LoadAllCSVsIbama("files")

	// Carregue o arquivo "Cadastro Básico" CSV na memória e gere o arquivo JSON
	//outputJSONFilePath := "resultEmpresas.json"
	//services.LoadCSVToMemory_CB_JSON("files/AILOSDB/BASE_AILOS.csv", outputJSONFilePath)

	//gera resulta empresa
	//services.LoadAll("files/AILOSDB/BASE_AILOS.csv")
	//
	//services.ExportJSON("resultEmpresas.json")

	//services.LoadAllCSVsICMBIO("files/INFRACOES_AMBIENTAIS/ICMBio")
	//services.LoadAllCSVsIbama("files/INFRACOES_AMBIENTAIS/IBAMA")
	//services.LoadCSVToMemory_Aneel()

	//services.ProcessaInfracoesAmbientaisIcmbio("files/AILOSDB/BASE_AILOS.csv")
	//services.ProcessaInfracoesAmbientaisIbama("files/AILOSDB/BASE_AILOS.csv")
	//services.ProcessaInfracoesAmbientaisAneel("files/AILOSDB/BASE_AILOS.csv")

	//services.LoadAllCSVs_SERPUB("files/INFRACOES_SERPUB")
	//services.ProcessarPEP("files/AILOSDB/BASE_AILOS.csv")

	//services.LoadAllCSVs_EMP_REG("files/EMP_ATIVIDADES_REG")
	//services.ProcessarAutorizacaoEspecial("files/AILOSDB/BASE_AILOS.csv")

	//services.LoadAllCSVs_Fraude("files/FRAUDE_CORRUPT")
	//services.ProcessarFraude("files/AILOSDB/BASE_AILOS.csv")

	//services.BuscaPF_CSV("files/INFRACOES_AMBIENTAIS/IBAMA", "files/AILOSDB/BASE_AILOS.csv")

	services.Convert("files/PF/cpsvm", "ultimosScrappsCpfs00h25m.json")

	//services.MatchPepsNames()
	log.Println("FINALIZOU")

	//log.Fatal(http.ListenAndServe(":8080", nil))
}
