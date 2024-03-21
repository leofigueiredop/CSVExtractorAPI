package services

import (
	"CSVExtractor/models"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
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
		if len(record) < 2 {
			log.Println("Invalid record:", record)
			continue
		}
		tipoPessoa := record[0]

		if tipoPessoa == "PJ" {
			cnpj := record[1]

			// Ensure cnpj is 8+ digits, pad with leading zeros if not
			cnpjFormatted := fmt.Sprintf("%014s", cnpj) // let's format CNPJ to have 14 characters, filling with zeros at left if needed

			cnpjBase := cnpjFormatted[:8]

			pessoa := models.Pessoa{
				Tipo_pessoa: tipoPessoa,
				CPF_CNPJ:    cnpjBase,
			}

			Pessoas[cnpjBase] = pessoa
		}
	}

	return nil
}

func LoadEmpresas() {
	files, err := os.ReadDir("files/PJ/empresas")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		err := loadEmpresa("files/PJ/empresas/" + file.Name())
		if err != nil {
			log.Fatalf("Error reading empresa CSV file "+file.Name()+": %s", err)
		}

		fmt.Println("Finished file " + file.Name())
	}
}

func LoadEstabelecimentos() {
	files, err := os.ReadDir("files/PJ/estabele")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		err := loadEstabelecimento("files/PJ/estabele/" + file.Name())
		if err != nil {
			log.Fatalf("Error reading estabele CSV file "+file.Name()+": %s", err)
		}

		fmt.Println("Finished file " + file.Name())
	}
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

		LoadEmpresas()

	}()

	go func() {
		defer wg.Done()
		LoadEstabelecimentos()

	}()

	wg.Wait()
	fmt.Println("All data loaded successfully")
}

func ExportJSON(outputFilePath string) {
	log.Println("Start exporting JSON data...")

	log.Printf("Number of pessoas: %d\n", len(Pessoas))

	// Abrindo arquivo de saída
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatalf("Could not create output file: %v", err)
	}
	defer outputFile.Close()

	// Configurando o encoder JSON
	encoder := json.NewEncoder(outputFile)
	encoder.SetIndent("", "  ")

	// Criando um slice para armazenar todas as pessoas
	allPessoas := make([]models.Pessoa, 0, len(Pessoas))

	// Iterando sobre cada pessoa
	for _, pessoa := range Pessoas {
		log.Printf("Exporting pessoa with CPF/CNPJ: %s\n", pessoa.CPF_CNPJ)

		// Associando empresas e estabelecimentos à pessoa
		empresas := Empresas[pessoa.CPF_CNPJ]
		for i := range empresas {
			empresa := &empresas[i]
			empresa.Estabelecimentos = Estabelecimentos[empresa.CnpjBase]
		}
		pessoa.Empresas = empresas

		// Adicionando a pessoa ao slice
		allPessoas = append(allPessoas, pessoa)
	}

	// Codificando e escrevendo todas as pessoas no arquivo de saída
	if err := encoder.Encode(allPessoas); err != nil {
		log.Printf("Could not encode pessoas: %v", err)
	}

	log.Println("Finished exporting JSON data.")
	log.Println("Data has been successfully written to", outputFilePath)
}
