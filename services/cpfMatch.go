package services

import (
	"CSVExtractor/models"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Registro struct {
	Cnpj      string                   `json:"cnpj"`
	Registros []models.RESULTAMBIENTAL `json:"registros"`
}

type RecordCsv struct {
	Nome string `json:"name"`
	Cpf  string `json:"cpf"`
}

func MatchCpfs() {

	dirResult := "PF/"
	dirCsv := "files/PF/cpfs/"

	var result []models.RESULTAMBIENTAL
	//var record RecordCsv

	CPFs := make(map[string]RecordCsv)
	Results := make(map[string]Registro)

	files, err := os.ReadDir(dirResult)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		file, err := os.ReadFile(dirResult + f.Name())
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(file, &result)

	}

	fmt.Println("Resultado: ", len(Results))

	files, err = os.ReadDir(dirCsv)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {

		fmt.Println("arquivo: ", f.Name())

		file, err := os.Open(dirCsv + f.Name())

		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		reader := csv.NewReader(file)
		reader.FieldsPerRecord = -1
		reader.Comma = ','
		reader.Read()
		for {
			fields, err := reader.Read()

			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("arquivo: ", f.Name())
				fmt.Println("erro: ", err)

			}
			if len(fields) > 3 {
				if fields[1] != "CPFValue" {
					CPFs[fields[1]] = RecordCsv{
						Cpf:  fields[1],
						Nome: fields[3],
					}
				}
			}
		}
	}
	fmt.Println("CPF: ", len(CPFs))

	// Verificar CPFs
	NotFound := []RecordCsv{}
	var FoundCounter int

	for _, cpf := range Results {
		if _, found := CPFs[cpf.Cnpj]; found {
			FoundCounter++
		} else {
			fmt.Println("NotFound: ", len(cpf.Registros))
			NotFound = append(NotFound, RecordCsv{
				Cpf:  cpf.Cnpj,
				Nome: cpf.Registros[0].NOME,
			})

		}
	}

	// Gravar arquivos não encontrados
	fileNotFound, _ := json.MarshalIndent(NotFound, "", " ")
	if err := os.WriteFile("notfound.json", fileNotFound, 0644); err != nil {
		panic(err)
	}

	fmt.Printf("CNPJ encontrados nos arquivos de CPFs: %d\nRegistros não encontrados: %d\n", FoundCounter, len(NotFound))
}

func MatchCpfsWithName() {

	dirResult := "resultPEPPF.json"
	dirCsv := "files/PF/cpfs/"

	var result []models.RESULTAMBIENTAL
	//var record RecordCsv

	CPFs := make(map[string]RecordCsv)
	Results := make(map[string]models.RESULTAMBIENTAL)

	file, err := os.ReadFile(dirResult)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(file, &result)

	for _, reg := range result {
		Results[reg.CPFCNPJ] = reg
	}

	fmt.Println("Resultado: ", len(Results))

	files, err := os.ReadDir(dirCsv)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {

		fmt.Println("arquivo: ", f.Name())

		file, err := os.Open(dirCsv + f.Name())

		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		reader := csv.NewReader(file)
		reader.FieldsPerRecord = -1
		reader.Comma = ','
		reader.Read()
		for {
			fields, err := reader.Read()

			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("arquivo: ", f.Name())
				fmt.Println("erro: ", err)

			}
			if len(fields) > 3 {
				if fields[1] != "CPFValue" {
					CPFs[fields[1]] = RecordCsv{
						Cpf:  fields[1],
						Nome: fields[3],
					}
				}
			}
		}
	}
	fmt.Println("CPF: ", len(CPFs))

	// Verificar CPFs
	NotFound := []models.RESULTAMBIENTAL{}
	Found := []models.RESULTAMBIENTAL{}
	NomeDifere := []models.RESULTAMBIENTAL{}

	for _, resultCpf := range Results {
		if recordCpf, found := CPFs[resultCpf.CPFCNPJ]; found {
			if strings.ToLower(resultCpf.NOME) == strings.ToLower(recordCpf.Nome) {
				fmt.Printf("Found: %s - %s - %s \n", recordCpf.Cpf, recordCpf.Nome, resultCpf.NOME)

				resultCpf.NomeAux = recordCpf.Nome
				Found = append(Found, resultCpf)
			} else {
				fmt.Printf("Difere: %s - %s - %s \n", recordCpf.Cpf, recordCpf.Nome, resultCpf.NOME)

				resultCpf.NomeAux = recordCpf.Nome
				NomeDifere = append(NomeDifere, resultCpf)
			}

		} else {

			fmt.Printf("Not Found: %s - %s - %s \n", recordCpf.Cpf, recordCpf.Nome, resultCpf.NOME)

			resultCpf.NomeAux = recordCpf.Nome
			NotFound = append(NotFound, resultCpf)

		}
	}

	// Gravar arquivos não encontrados
	fileNotFound, _ := json.MarshalIndent(NotFound, "", " ")
	if err := os.WriteFile("notfound.json", fileNotFound, 0644); err != nil {
		panic(err)
	}
	fileFound, _ := json.MarshalIndent(Found, "", " ")
	if err := os.WriteFile("found.json", fileFound, 0644); err != nil {
		panic(err)
	}
	diferente, _ := json.MarshalIndent(NomeDifere, "", " ")
	if err := os.WriteFile("diferent.json", diferente, 0644); err != nil {
		panic(err)
	}
	fmt.Printf("CNPJ encontrados nos arquivos de CPFs: %d\nRegistros não encontrados: %d\n", len(Found), len(NotFound))
}
