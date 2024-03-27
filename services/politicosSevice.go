package services

import (
	"CSVExtractor/models"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/text/encoding/charmap"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

var (
	pepData       = make(map[string][]models.PEP)
	expulsoesData = make(map[string][]Expulso)
	sancoesData   = make(map[string][]Sancao)
	candidatos    = make(map[string][]Candidato)
)

// Expulso represents the data in the CSV file you provided
type Expulso struct {
	Cadastro               string
	CodigoDaSancao         string
	TipoDePessoa           string
	CPF                    string
	NomeDoSancionado       string
	CategoriaDaSancao      string
	NumeroDoDocumento      string
	NumeroDoProcesso       string
	DataInicioSancao       string
	DataFinalSancao        string
	DataPublicacao         string
	Publicacao             string
	Detalhamento           string
	DataDoTransito         string
	Abragencia             string
	CargoEfetivo           string
	FuncaoOuCargo          string
	OrgaoDeLotacao         string
	OrgaoSancionador       string
	UFOrgaoSancionador     string
	EsferaOrgaoSancionador string
	FundamentacaoLegal     string
}

// LoadCSVToMemory_EXP loads the CSV file to memory
func LoadCSVToMemory_EXP(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read input file: %s", filePath, err)
	}
	defer file.Close()
	r := charmap.ISO8859_1.NewDecoder().Reader(file)
	csvReader := csv.NewReader(r)
	csvReader.Comma = ';'
	csvReader.LazyQuotes = true

	re := regexp.MustCompile("[^0-9\\*]+")

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error reading CSV file: %s", err)
		}

		cleanCPF := re.ReplaceAllString(record[3], "")

		expulso := Expulso{
			Cadastro:               strings.Trim(record[0], "\""),
			CodigoDaSancao:         strings.Trim(record[1], "\""),
			TipoDePessoa:           strings.Trim(record[2], "\""),
			CPF:                    cleanCPF,
			NomeDoSancionado:       strings.Trim(record[4], "\""),
			CategoriaDaSancao:      strings.Trim(record[5], "\""),
			NumeroDoDocumento:      strings.Trim(record[6], "\""),
			NumeroDoProcesso:       strings.Trim(record[7], "\""),
			DataInicioSancao:       strings.Trim(record[8], "\""),
			DataFinalSancao:        strings.Trim(record[9], "\""),
			DataPublicacao:         strings.Trim(record[10], "\""),
			Publicacao:             strings.Trim(record[11], "\""),
			Detalhamento:           strings.Trim(record[12], "\""),
			DataDoTransito:         strings.Trim(record[13], "\""),
			Abragencia:             strings.Trim(record[14], "\""),
			CargoEfetivo:           strings.Trim(record[15], "\""),
			FuncaoOuCargo:          strings.Trim(record[16], "\""),
			OrgaoDeLotacao:         strings.Trim(record[17], "\""),
			OrgaoSancionador:       strings.Trim(record[18], "\""),
			UFOrgaoSancionador:     strings.Trim(record[19], "\""),
			EsferaOrgaoSancionador: strings.Trim(record[20], "\""),
			FundamentacaoLegal:     strings.Trim(record[21], "\""),
		}

		log.Println("Loading to memory Expulsoes", cleanCPF)
		expulsoesData[cleanCPF] = append(expulsoesData[cleanCPF], expulso)
	}
}

func LoadCSVToMemory_PEP(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read input file: %s", err)
	}
	defer file.Close()
	r := charmap.ISO8859_1.NewDecoder().Reader(file)
	csvReader := csv.NewReader(r)
	csvReader.Comma = ';'
	csvReader.LazyQuotes = true
	csvReader.FieldsPerRecord = -1

	// Expressão regular para manter os asteriscos e numeros do CPF
	re := regexp.MustCompile("[^0-9\\*]+")

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading CSV file: %s", err)
		}

		cleanCPF := re.ReplaceAllString(record[0], "")

		pep := models.PEP{
			UUID:                  uuid.NewString(),
			CPF:                   cleanCPF, // Alterado para cleanCPF
			Nome_PEP:              strings.Trim(record[1], "\""),
			Sigla_Funcao:          strings.Trim(record[2], "\""),
			Descricao_Funcao:      strings.Trim(record[3], "\""),
			Nivel_Funcao:          strings.Trim(record[4], "\""),
			Nome_Orgao:            strings.Trim(record[5], "\""),
			Data_Inicio_Exercicio: strings.Trim(record[6], "\""),
			Data_Fim_Exercicio:    strings.Trim(record[7], "\""),
			Data_Fim_Carencia:     strings.Trim(record[8], "\""),
		}

		log.Println("Loading to memory PEP", cleanCPF)
		pepData[cleanCPF] = append(pepData[cleanCPF], pep)
	}
}

type Sancao struct {
	CPFCNPJ              string
	NomeSancionado       string
	Cadastro             string
	UFSancionado         string
	NomeOrgaoSancionador string
	CategoriaSancao      string
	DataPublicacao       string
	ValorMulta           string
	Quantidade           string
}

func LoadCSVToMemory_SANCAO(filePath string) map[string][]Sancao {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read input file: %s", filePath, err)
	}
	defer file.Close()
	r := charmap.ISO8859_1.NewDecoder().Reader(file)
	csvReader := csv.NewReader(r)
	csvReader.Comma = ';'
	csvReader.LazyQuotes = true

	re := regexp.MustCompile("[^0-9\\*]+")
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading CSV file: %s", err)
		}
		cleanCPF := re.ReplaceAllString(record[0], "")

		sancao := Sancao{
			CPFCNPJ:              cleanCPF,
			NomeSancionado:       strings.Trim(record[1], "\""),
			Cadastro:             strings.Trim(record[2], "\""),
			UFSancionado:         strings.Trim(record[3], "\""),
			NomeOrgaoSancionador: strings.Trim(record[4], "\""),
			CategoriaSancao:      strings.Trim(record[5], "\""),
			DataPublicacao:       strings.Trim(record[6], "\""),
			ValorMulta:           strings.Trim(record[7], "\""),
			Quantidade:           strings.Trim(record[8], "\""),
		}

		log.Println("Loading to memory Sancoes", cleanCPF)
		sancoesData[cleanCPF] = append(sancoesData[cleanCPF], sancao)
	}
	return sancoesData
}

func ofuscarCPF(cpfCompleto string) string {
	// Verifique se o CPF completo tem a forma correta
	re := regexp.MustCompile(`^(\d{3})(\d{3})(\d{3})(\d{2})$`)
	if re.MatchString(cpfCompleto) {
		return "***" + cpfCompleto[3:6] + cpfCompleto[6:9] + "**"
	} else {
		return ""
	}
}

type Candidato struct {
	Nome       string `json:"Nome"`
	CPF        string `json:"CPF"`
	AnoEleicao string `json:"ANO_ELEICAO"`
	SgUF       string `json:"SG_UF"`
	Cargo      string `json:"DS_CARGO"`
}

func LoadCSVToMemory_CANDIDATOS() {
	imputDir := "./files/INFRACOES_SERPUB/consulta_cand_2022/"

	files, err := os.ReadDir(imputDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".csv" {
			filePath := filepath.Join(imputDir, file.Name())
			file, err := os.Open(filePath)
			if err != nil {
				log.Println("Unable to read input file: %s", err)
			}
			log.Println("reading ", file.Name())
			defer file.Close()

			r := charmap.ISO8859_1.NewDecoder().Reader(file)
			csvReader := csv.NewReader(r)
			csvReader.Comma = ';'
			csvReader.LazyQuotes = true
			csvReader.FieldsPerRecord = -1

			// Expressão regular para manter os asteriscos e numeros do CPF
			re := regexp.MustCompile("[^0-9]+")
			csvReader.Read()
			for {
				record, err := csvReader.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatalf("Error reading CSV file: %s", err)
				}
				cleanCPF := re.ReplaceAllString(record[20], "")

				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatal(err)
				}

				data := Candidato{
					Nome:       record[18],
					CPF:        cleanCPF,
					AnoEleicao: record[3],
					SgUF:       record[11],
					Cargo:      record[15],
				}
				log.Println("add candidato: %s", cleanCPF)
				candidatos[cleanCPF] = append(candidatos[cleanCPF], data)

			}
		}
	}
}

func ProcessarPEP(filePath string) {

	LoadCSVToMemory_CANDIDATOS()

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read input file: %s", err)
	}
	defer file.Close()
	r := charmap.ISO8859_1.NewDecoder().Reader(file)
	reader := csv.NewReader(r)
	reader.Comma = ';'
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1

	isFirstRow := true
	var finalCompanies []models.RESULTAMBIENTAL
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
		if tipoPessoa == "PF" {
			cnpj := record[1]
			cnpjBase := fmt.Sprintf("%011s", cnpj)
			subCNPJBase := ofuscarCPF(cnpjBase)

			if len(candidatos[cnpjBase]) > 0 {

				// Process sanctionsData
				if records, ok := pepData[subCNPJBase]; ok {
					for _, record := range records {
						r := models.RESULTAMBIENTAL{
							DATASET:     "SERV_PUB",
							TIPO:        "PEP",
							CPFCNPJ:     cnpj,
							CODCONTROLE: candidatos[cnpjBase][0].CPF,
							NOME:        record.Nome_PEP,
							Data:        record.Data_Inicio_Exercicio,
							NumProcesso: "",
							Descricao:   fmt.Sprintf("Função: %s  - Orgão: %s  - Data início exercicio: %s - Data fim carência: %s ", record.Descricao_Funcao, record.Nome_Orgao, record.Data_Inicio_Exercicio, record.Data_Fim_Carencia),
						}
						finalCompanies = append(finalCompanies, r)
					}
				}

				// Process expulsosData
				if records, ok := expulsoesData[subCNPJBase]; ok {
					for _, record := range records {
						r := models.RESULTAMBIENTAL{
							DATASET:     "SERV_PUB",
							TIPO:        "EXPULSO",
							CPFCNPJ:     cnpj,
							NOME:        record.NomeDoSancionado,
							CODCONTROLE: record.Cadastro,
							Data:        record.DataDoTransito,
							NumProcesso: record.NumeroDoProcesso,
							Descricao:   fmt.Sprintf("Função: %s  - Orgão: %s  - Detalhes: %s -  %s ", record.FuncaoOuCargo, record.OrgaoSancionador, record.Detalhamento, record.FundamentacaoLegal),
						}
						finalCompanies = append(finalCompanies, r)
					}
				}

				// Process sancoesData
				if records, ok := sancoesData[subCNPJBase]; ok {
					for _, record := range records {
						r := models.RESULTAMBIENTAL{
							DATASET:     "SERV_PUB",
							TIPO:        "SANSAO",
							CPFCNPJ:     cnpj,
							CODCONTROLE: record.Cadastro,
							NOME:        record.NomeSancionado,
							Data:        record.DataPublicacao,
							NumProcesso: "",
							Descricao:   fmt.Sprintf("Orgão: %s  - Multa: %s  - Detalhes: %s", record.NomeOrgaoSancionador, record.ValorMulta, record.CategoriaSancao),
						}
						finalCompanies = append(finalCompanies, r)
					}
				}

			}
		}
	}

	// create JSON from finalCompanies slice
	jsonData, err := json.MarshalIndent(finalCompanies, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	// write JSON data to a file only if there is data to write
	if len(finalCompanies) > 0 {
		err = os.WriteFile("resultPEPPFCandidatos.json", jsonData, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func LoadAllCSVs_SERPUB(dirPath string) {

	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".csv" {
			wg.Add(1)
			go func(file os.DirEntry) {
				defer wg.Done()
				filePath := filepath.Join(dirPath, file.Name())
				switch {
				case strings.Contains(strings.ToLower(file.Name()), "pep"):
					LoadCSVToMemory_PEP(filePath)
				case strings.Contains(strings.ToLower(file.Name()), "expulsoes"):
					LoadCSVToMemory_EXP(filePath)
				case strings.Contains(strings.ToLower(file.Name()), "sancoes"): // adicionado este case para lidar com arquivos de Sancao
					LoadCSVToMemory_SANCAO(filePath)
				default:
					log.Printf("No suitable function for the file %s", filePath)
				}
			}(file)
		}
	}
	wg.Wait()
}
