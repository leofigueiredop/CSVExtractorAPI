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
	"sync"
)

type ANP struct {
	StatusProcesso      string `csv:"Status Processo"`
	Superintendencia    string `csv:"Superintendência"`
	NumeroDoProcesso    string `csv:"Número do Processo"`
	NumeroDoDUF         string `csv:"Número do DUF"`
	CNPJCPF             string `csv:"CNPJ/CPF"`
	RazaoSocial         string `csv:"Razão Social"`
	DataTransitoJulgado string `csv:"Data Transito Julgado"`
	Vencimento          string `csv:"Vencimento"`
	ValorDaMulta        string `csv:"Valor da Multa"`
	ValorTotalPago      string `csv:"Valor Total Pago"`
}
type empresaAnpAtt struct {
	CNPJ      string                   `json:"cnpj"`
	Registros []models.RESULTAMBIENTAL `json:"registros"`
}
type AttSancoes struct {
	CPFCNPJ    string `csv:"CNPJ/CPF da Pessoa ou Empresa Sancionada"`
	Nome       string `csv:"Nome da Pessoa ou Empresa Sancionada"`
	Cadastro   string `csv:"Cadastro"`
	UF         string `csv:"UF sancionado"`
	Orgao      string `csv:"Nome do Órgão Sancionador"`
	Categoria  string `csv:"Categoria sanção"`
	Data       string `csv:"Data Publicação"`
	ValorMulta string `csv:"Valor multa"`
	Quantidade string `csv:"Quantidade"`
}

var (
	anpData        = make(map[string][]ANP)
	attData        = make(map[string][]ATT)
	attSancoesData = make(map[string][]AttSancoes)
)

var empresasANtAtt []empresaAnpAtt

func LoadCSVToMemory_ANP(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file ", filePath, err)
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
		} else if err != nil {
			log.Printf("Error reading ANP CSV file: %s", err)
			continue
		}
		// Check for valid number of columns
		if len(record) < 10 {
			log.Println("Skipping record due to insufficient columns:", record)
			continue
		}
		cleanCPF := re.ReplaceAllString(record[4], "")
		anp := ANP{
			StatusProcesso:      strings.Trim(record[0], "\""),
			Superintendencia:    strings.Trim(record[1], "\""),
			NumeroDoProcesso:    strings.Trim(record[2], "\""),
			NumeroDoDUF:         strings.Trim(record[3], "\""),
			CNPJCPF:             cleanCPF,
			RazaoSocial:         strings.Trim(record[5], "\""),
			DataTransitoJulgado: strings.Trim(record[6], "\""),
			Vencimento:          strings.Trim(record[7], "\""),
			ValorDaMulta:        strings.Trim(record[8], "\""),
			ValorTotalPago:      strings.Trim(record[9], "\""),
		}

		log.Println("Loading to memory ANP Record ", anp.NumeroDoProcesso)
		anpData[cleanCPF] = append(anpData[cleanCPF], anp)
	}
}

type ATT struct {
	Autuado string `csv:"AUTUADO"`
	CPFCNPJ string `csv:"CPF/CNPJ"`
}

func LoadCSVToMemory_ATT(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read ATT input file: %s", err)
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
			log.Fatalf("Error reading ATT CSV file: %s", err)
			return
		}

		// Only keeps numerical characters
		cleanCPFCNPJ := re.ReplaceAllString(record[1], "")

		attRecord := ATT{
			Autuado: strings.Trim(record[0], "\""),
			CPFCNPJ: cleanCPFCNPJ,
		}

		log.Printf("Loading to memory ATT Record with CPFCNPJ %s\n", attRecord.CPFCNPJ)
		attData[attRecord.CPFCNPJ] = append(attData[attRecord.CPFCNPJ], attRecord)
	}
	log.Println("Loading to memory ATT has complete.")
}

func LoadCSVToMemory_AttSancoes(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read ATT Sancoes input file: %s", err)
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
			log.Fatalf("Error reading ATT Sancoes CSV file: %s", err)
			return
		}

		// Only keeps numerical characters
		cleanCPFCNPJ := re.ReplaceAllString(record[0], "")
		for i, val := range record {
			record[i] = strings.Trim(val, "\"")
		}
		attSancoesRecord := AttSancoes{
			CPFCNPJ:    cleanCPFCNPJ,
			Nome:       strings.Trim(record[1], "\""),
			Cadastro:   strings.Trim(record[2], "\""),
			UF:         strings.Trim(record[3], "\""),
			Orgao:      strings.Trim(record[4], "\""),
			Categoria:  strings.Trim(record[5], "\""),
			Data:       strings.Trim(record[6], "\""),
			ValorMulta: strings.Trim(record[7], "\""),
			Quantidade: strings.Trim(record[8], "\""),
		}

		log.Printf("Loading to memory ATT Sancoes Record with CNPJ/CPF %s\n", attSancoesRecord.CPFCNPJ)
		attSancoesData[attSancoesRecord.CPFCNPJ] = append(attSancoesData[attSancoesRecord.CPFCNPJ], attSancoesRecord)
	}
	log.Println("Loading to memory AttSancoes has complete.")
}

func ProcessaAnpAtt(filePath string) {
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
	var finalCompanies []empresaAnpAtt
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
			emp := empresaAnpAtt{CNPJ: cnpjBase}

			if records, ok := anpData[cnpjBase]; ok {
				for _, record := range records {
					r := models.RESULTAMBIENTAL{
						DATASET:     "ANP",
						TIPO:        "AUTUACAO",
						CPFCNPJ:     cnpjBase,
						CODCONTROLE: record.NumeroDoProcesso,
						Data:        record.DataTransitoJulgado,
						NumProcesso: record.NumeroDoProcesso,
						Descricao:   fmt.Sprintf("%s - Valor da Multa: %s - %s", record.StatusProcesso, record.ValorDaMulta, record.Superintendencia),
					}
					emp.Registros = append(emp.Registros, r)
				}
			}
			// Populate the data from attData
			if records, ok := attData[cnpjBase]; ok {
				for _, record := range records {
					r := models.RESULTAMBIENTAL{
						DATASET:     "ATT",
						TIPO:        "AUTUACAO", // Update this if needed
						CPFCNPJ:     cnpjBase,
						NOME:        record.Autuado, // Update this if needed
						CODCONTROLE: "N/A",          // Update this if needed
					}
					emp.Registros = append(emp.Registros, r)
				}
			}
			// Populate the data from attSancoesData
			if records, ok := attSancoesData[cnpjBase]; ok {
				for _, record := range records {
					r := models.RESULTAMBIENTAL{
						CODCONTROLE: record.Cadastro,
						DATASET:     "ATT",
						TIPO:        "SANSAO",
						CPFCNPJ:     cnpjBase,
						NOME:        record.Nome,
						Data:        record.Data,
						Descricao:   fmt.Sprintf("Categoria: %s - Valor da Multa: %s - %s", record.Categoria, record.ValorMulta, record.Orgao),
					}
					emp.Registros = append(emp.Registros, r)
				}
			}
			finalCompanies = append(finalCompanies, emp)
		}
	}

	// Filter empresas slice for companies that have at least one record
	var empresasWithRecords []empresaAnpAtt
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
		err = os.WriteFile("resultAnpAtt.json", jsonData, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func LoadAllCSVs(dirPath string) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	for _, file := range files {
		fileName := strings.ToLower(file.Name())
		filePath := filepath.Join(dirPath, file.Name())
		if filepath.Ext(file.Name()) == ".csv" {
			if strings.Contains(fileName, "multas") {
				LoadCSVToMemory_ANP(filePath)
			} else {

				wg.Add(1)
				go func(file os.DirEntry) {
					defer wg.Done()

					switch {
					case strings.Contains(fileName, "multas"):
						LoadCSVToMemory_ANP(filePath)
					case strings.Contains(fileName, "att"):
						LoadCSVToMemory_ATT(filePath)
					case strings.Contains(fileName, "sancoes"):
						LoadCSVToMemory_AttSancoes(filePath)
					default:
						log.Printf("No suitable function for the file %s", filePath)
					}
				}(file)
			}
		}
	}
	wg.Wait()
}
