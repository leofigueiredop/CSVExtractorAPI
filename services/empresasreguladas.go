package services

import (
	"CSVExtractor/models"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type AutorizacaoEspecial struct {
	CNPJ                       string `csv:"CNPJ"`
	RazaoSocial                string `csv:"Razão Social"`
	CodigoCategoria            string `csv:"Código da categoria"`
	DescricaoCategoria         string `csv:"Descrição da categoria"`
	CodigoAtividade            string `csv:"Código da atividade"`
	DescricaoAtividade         string `csv:"Descrição da atividade"`
	DataInicioAtividade        string `csv:"Data de início da atividade"`
	DataTerminoAtividade       string `csv:"Data de término da atividade"`
	PotencialPoluidorAtividade string `csv:"Potenc. de Pol. da atividade"`
	Municipio                  string `csv:"Município"`
	Estado                     string `csv:"Estado"`
	Latitude                   string `csv:"Latitude"`
	Longitude                  string `csv:"Longitude"`
	SituacaoCadastral          string `csv:"Situação cadastral"`
	UltimaAtualizacao          string `csv:"Última Atualização"`
}

type empresaAutorizacaoEspecial struct {
	CNPJ      string                   `json:"cnpj"`
	Registros []models.RESULTAMBIENTAL `json:"registros"`
}

var (
	anpDataAutorizacaoEspecial = make(map[string][]AutorizacaoEspecial)
)

func LoadCSVToMemory_AutorizacaoEspecial(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read input file: %s", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'
	csvReader.Read()

	// Regular expression to remove all non-numeric characters
	re := regexp.MustCompile("[^0-9]+")

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading AutorizacaoEspecial CSV file: %s", err)
			return
		}

		// Only keeps numerical characters
		cleanCPFCNPJ := re.ReplaceAllString(record[0], "")

		for i, val := range record {
			record[i] = strings.Trim(val, "\"")
		}

		autorizacaoEspecialRecord := AutorizacaoEspecial{
			CNPJ:                       cleanCPFCNPJ,
			RazaoSocial:                record[1],
			CodigoCategoria:            record[2],
			DescricaoCategoria:         record[3],
			CodigoAtividade:            record[4],
			DescricaoAtividade:         record[5],
			DataInicioAtividade:        record[6],
			DataTerminoAtividade:       record[7],
			PotencialPoluidorAtividade: record[8],
			Municipio:                  record[9],
			Estado:                     record[10],
			Latitude:                   record[11],
			Longitude:                  record[12],
			SituacaoCadastral:          record[13],
			UltimaAtualizacao:          record[14],
		}

		log.Printf("Loading to memory AutorizacaoEspecial Record with CNPJ/CPF %s\n", autorizacaoEspecialRecord.CNPJ)
		anpDataAutorizacaoEspecial[autorizacaoEspecialRecord.CNPJ] = append(anpDataAutorizacaoEspecial[autorizacaoEspecialRecord.CNPJ], autorizacaoEspecialRecord)
	}
	log.Println("Loading to memory AutorizacaoEspecial has complete.")
}

func ProcessarAutorizacaoEspecial(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Unable to read input file ", filePath, ". Error: ", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1

	isFirstRow := true
	var finalCompanies []empresaAutorizacaoEspecial
	for {
		// Read each record from csv
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error reading the file records: ", err)
			continue
		}
		// Skip the header
		if isFirstRow {
			isFirstRow = false
			continue
		}
		tipoPessoa := record[0]
		if tipoPessoa == "PJ" {
			cnpj := record[1] // Modify this index according to your csv structure
			cnpjBase := fmt.Sprintf("%014s", cnpj)
			emp := empresaAutorizacaoEspecial{CNPJ: cnpjBase}

			if records, ok := anpDataAutorizacaoEspecial[cnpjBase]; ok {
				for _, record := range records {
					r := models.RESULTAMBIENTAL{
						DATASET:     "EMPRESA_REGULADA",
						TIPO:        "AUTORIZACAO",
						CPFCNPJ:     cnpjBase,
						CODCONTROLE: record.SituacaoCadastral,
						Data:        record.DataInicioAtividade,
						NumProcesso: "",
						Descricao:   fmt.Sprintf("Data fim Atividade: %s - Ultima Atualicacao: %s - Categoria: %s - Atividade: %s - Potencial Poluidor: %s", record.DataTerminoAtividade, record.UltimaAtualizacao, record.DescricaoCategoria, record.DescricaoAtividade, record.PotencialPoluidorAtividade),
					}
					emp.Registros = append(emp.Registros, r)
				}
			}
			finalCompanies = append(finalCompanies, emp)
		}
	}

	// Filter empresas slice for companies that have at least one record
	var empresasWithRecords []empresaAutorizacaoEspecial
	for _, emp := range finalCompanies {
		if len(emp.Registros) > 0 {
			empresasWithRecords = append(empresasWithRecords, emp)
		}
	}

	// create JSON from the empresasWithRecords slice
	jsonData, err := json.MarshalIndent(empresasWithRecords, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	// write JSON data to a file only if there is data to write
	if len(empresasWithRecords) > 0 {
		err = os.WriteFile("resultAutorizacaoEspecial.json", jsonData, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func LoadAllCSVs_EMP_REG(dirPath string) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".csv" {
			filePath := filepath.Join(dirPath, file.Name())
			LoadCSVToMemory_AutorizacaoEspecial(filePath)
		}
	}
}
