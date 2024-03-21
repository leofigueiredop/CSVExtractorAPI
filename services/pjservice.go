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
	"strings"
	"sync"
)

var (
	Pessoas          = make(map[string]models.Pessoa)            // map with key: CPF/CNPJ
	Empresas         = make(map[string][]models.Empresa)         // map with key: CNPJ Base
	Estabelecimentos = make(map[string][]models.Estabelecimento) // map with key: CNPJ Base
)

func LoadPessoas(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file ", filePath, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1

	isFirstRow := true
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Skip the header
		if isFirstRow {
			isFirstRow = false
			continue
		}

		cnpj := record[1][:8] // CNPJ base is the first 8 characters
		pessoa := models.Pessoa{
			Tipo_pessoa: record[0],
			CPF_CNPJ:    cnpj,
		}

		Pessoas[cnpj] = pessoa
	}

	return nil
}

func LoadEmpresas() error {
	return filepath.Walk("PJ/", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if strings.HasSuffix(path, ".EMPRECSV") {
				err := loadEmpresa(path)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func LoadEstabelecimentos() error {
	return filepath.Walk("PJ/", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if strings.HasSuffix(path, ".ESTABELE") {
				err := loadEstabelecimento(path)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func loadEmpresa(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file ", filePath, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	reader.FieldsPerRecord = -1

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("error reading record, skipping: ", err)
			continue
		}

		cnpjBase := record[0]
		if _, exists := Pessoas[cnpjBase]; exists {
			Empresas[cnpjBase] = append(Empresas[cnpjBase], models.Empresa{
				CnpjBase:       cnpjBase,
				RazaoSocial:    record[1],
				Natureza:       record[2],
				Qualificacao:   record[3],
				CapitalSocial:  record[4],
				Porte:          record[5],
				EnteFederativo: record[6],
			})
		}
	}

	return nil
}

func loadEstabelecimento(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file ", filePath, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	reader.FieldsPerRecord = -1

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("error reading record, skipping: ", err)
			continue
		}

		cnpjBase := record[0]
		if _, exists := Pessoas[cnpjBase]; exists {
			Estabelecimentos[cnpjBase] = append(Estabelecimentos[cnpjBase], models.Estabelecimento{
				CnpjBase:             record[0],
				CnpjOrdem:            record[1],
				CnpjDv:               record[2],
				IdentificadorMatriz:  record[3],
				NomeFantasia:         record[4],
				SituacaoCadastral:    record[5],
				DataSituacao:         record[6],
				MotivoSituacao:       record[7],
				CidadeExterior:       record[8],
				Pais:                 record[9],
				DataInicioAtividade:  record[10],
				CnaeFiscalPrincipal:  record[11],
				CnaeFiscalSecundaria: record[12],
				TipoLogradouro:       record[13],
				Logradouro:           record[14],
				Numero:               record[15],
				Complemento:          record[16],
				Bairro:               record[17],
				Cep:                  record[18],
				Uf:                   record[19],
				Municipio:            record[20],
				Ddd1:                 record[21],
				Telefone1:            record[22],
				Ddd2:                 record[23],
				Telefone2:            record[24],
				DddFax:               record[25],
				Fax:                  record[26],
				Email:                record[27],
				SituacaoEspecial:     record[28],
				DataSituacaoEspecial: record[29],
			})
		}
	}

	return nil
}

func LoadAll(filePath string) {
	err := LoadPessoas(filePath)
	if err != nil {
		log.Fatal("Unable to load pessoas: ", err)
	}

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		err = LoadEmpresas()
		if err != nil {
			log.Fatal("Unable to load empresas: ", err)
		}
	}()

	go func() {
		defer wg.Done()
		err = LoadEstabelecimentos()
		if err != nil {
			log.Fatal("Unable to load estabelecimentos: ", err)
		}
	}()

	wg.Wait()
	fmt.Println("All data loaded successfully")
}

func ExportJSON(outputFilePath string) {
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatalf("Could not create output file: %v", err)
	}
	defer outputFile.Close()

	encoder := json.NewEncoder(outputFile)
	encoder.SetIndent("", "  ")

	for _, pessoa := range Pessoas {
		if err := encoder.Encode(pessoa); err != nil {
			log.Printf("Could not encode persona: %v, err: %v", pessoa, err)
		}
	}

	fmt.Println("Data has been successfully written to", outputFilePath)
}
