package services

import (
	"CSVExtractor/models"
	"encoding/csv"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

var (
	pepData                 = make(map[string][]models.PEP)
	ceisData                = make(map[string][]models.CEIS)
	cnepData                = make(map[string][]models.CNEP)
	autosInfracaoIbamaData  = make(map[string][]models.AutosInfracaoIbama)
	autosInfracaoICMBIOData = make(map[string][]models.AutosInfracaoICMBIO)
	trabalhoEscravoData     = make(map[string][]models.TrabalhoEscravo)
	suspensaobamaData       = make(map[string][]models.Suspensaobama)
	apreensaoIbamaData      = make(map[string][]models.ApreensaoIbama)
)

func LoadCSVToMemory_PEP(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read input file: %s", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'

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

		log.Println("Loading to memory PEP", pep.CPF)
		pepData[cleanCPF] = append(pepData[cleanCPF], pep)
	}
}

func LoadCSVToMemory_CNEP(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read CNEP input file: %s", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'

	// Expressão regular para remover todos os caracteres não numéricos
	re := regexp.MustCompile("[^0-9]+")

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading CNEP CSV file: %s", err)
			return
		}

		cleanCPF := re.ReplaceAllString(record[3], "")

		cnep := models.CNEP{
			UUID:                          uuid.NewString(),
			Cadastro:                      strings.Trim(record[0], "\""),
			CodigoSancao:                  strings.Trim(record[1], "\""),
			CPFCNPJSancionado:             cleanCPF,
			NomeSancionado:                strings.Trim(record[4], "\""),
			NomeInformadoOrgaoSancionador: strings.Trim(record[5], "\""),
			RazaoSocialCadastroReceita:    strings.Trim(record[6], "\""),
			NomeFantasiaCadastroReceita:   strings.Trim(record[7], "\""),
			NumeroProcesso:                strings.Trim(record[8], "\""),
			CategoriaSancao:               strings.Trim(record[9], "\""),
			ValorMulta:                    strings.Trim(record[10], "\""),
			DataInicioSancao:              strings.Trim(record[11], "\""),
			DataFinalSancao:               strings.Trim(record[12], "\""),
			DataPublicacao:                strings.Trim(record[13], "\""),
			Publicacao:                    strings.Trim(record[14], "\""),
			Detalhamento:                  strings.Trim(record[15], "\""),
			DataTransitoJulgado:           strings.Trim(record[16], "\""),
			AbrangenciaDecisaoJudicial:    strings.Trim(record[17], "\""),
			OrgaoSancionador:              strings.Trim(record[18], "\""),
			UfOrgaoSancionador:            strings.Trim(record[19], "\""),
			FundamentacaoLegal:            strings.Trim(record[21], "\""),
		}

		log.Println("Loading to memory CNEP", cnep.CPFCNPJSancionado)
		cnepData[cleanCPF] = append(cnepData[cleanCPF], cnep)
	}
}

func LoadCSVToMemory_CEIS(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read input file: %s", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'

	// Expressão regular para remover todos os caracteres não numéricos
	re := regexp.MustCompile("[^0-9]+")

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading CSV file: %s", err)
		}

		cleanCPF := re.ReplaceAllString(record[3], "")

		ceis := models.CEIS{
			UUID:                          uuid.NewString(),
			Cadastro:                      strings.Trim(record[0], "\""),
			CodigoSancao:                  strings.Trim(record[1], "\""),
			CPFCNPJSancionado:             cleanCPF,
			NomeSancionado:                strings.Trim(record[4], "\""),
			NomeInformadoOrgaoSancionador: strings.Trim(record[5], "\""),
			RazaoSocialCadastroReceita:    strings.Trim(record[6], "\""),
			NomeFantasiaCadastroReceita:   strings.Trim(record[7], "\""),
			NumeroProcesso:                strings.Trim(record[8], "\""),
			CategoriaSancao:               strings.Trim(record[9], "\""),
			DataInicioSancao:              strings.Trim(record[10], "\""),
			DataFinalSancao:               strings.Trim(record[11], "\""),
			DataPublicacao:                strings.Trim(record[12], "\""),
			Publicacao:                    strings.Trim(record[13], "\""),
			Detalhamento:                  strings.Trim(record[14], "\""),
			DataTransitoJulgado:           strings.Trim(record[15], "\""),
			AbrangenciaDecisaoJudicial:    strings.Trim(record[16], "\""),
			OrgaoSancionador:              strings.Trim(record[17], "\""),
			UfOrgaoSancionador:            strings.Trim(record[18], "\""),
			FundamentacaoLegal:            strings.Trim(record[20], "\""),
		}

		log.Println("Loading to memory CEIS", ceis.CPFCNPJSancionado)
		ceisData[cleanCPF] = append(ceisData[cleanCPF], ceis)
	}
}

func LoadCSVToMemory_AutosInfracaoIbama(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read AutosInfracaoIbama input file: %s", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'

	// Expressão regular para remover todos os caracteres não numéricos
	re := regexp.MustCompile("[^0-9]+")

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error reading AutosInfracaoIbama CSV file: %s", err)
			continue
		}

		cleanCPF := re.ReplaceAllString(record[21], "")

		autoInfracao := models.AutosInfracaoIbama{
			UUID:                uuid.NewString(),
			SeqAutoInfracao:     strings.Trim(record[0], "\""),
			NumAutoInfracao:     strings.Trim(record[1], "\""),
			SerAutoInfracao:     strings.Trim(record[2], "\""),
			TipoAuto:            strings.Trim(record[3], "\""),
			TipoMulta:           strings.Trim(record[4], "\""),
			ValAutoInfracao:     strings.Trim(record[5], "\""),
			PatrimonioApuracao:  strings.Trim(record[6], "\""),
			GravidadeInfracao:   strings.Trim(record[7], "\""),
			UnidArrecadacao:     strings.Trim(record[8], "\""),
			DesAutoInfracao:     strings.Trim(record[9], "\""),
			DatHoraAutoInfracao: strings.Trim(record[10], "\""),
			DatCienciaAutuacao:  strings.Trim(record[12], "\""),
			Municipio:           strings.Trim(record[14], "\""),
			Uf:                  strings.Trim(record[15], "\""),
			NumProcesso:         strings.Trim(record[16], "\""),
			CodInfracao:         strings.Trim(record[17], "\""),
			DesInfracao:         strings.Trim(record[18], "\""),
			TipoInfracao:        strings.Trim(record[19], "\""),
			NomeInfrator:        strings.Trim(record[20], "\""),
			CpfCnpjInfrator:     cleanCPF,
			DesLocalInfracao:    strings.Trim(record[28], "\""),
			TipoAcao:            strings.Trim(record[32], "\""),
			Operacao:            strings.Trim(record[33], "\""),
			DatLancamento:       strings.Trim(record[38], "\""),
		}

		log.Println("Loading to memory AutosInfracaoIbama ", autoInfracao.SeqAutoInfracao)
		autosInfracaoIbamaData[cleanCPF] = append(autosInfracaoIbamaData[cleanCPF], autoInfracao)
	}
}

func LoadCSVToMemory_AutosInfracaoICMBIO(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read AutosInfracaoICMBIO input file: %s", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'

	// Expressão regular para remover todos os caracteres não numéricos
	re := regexp.MustCompile("[^0-9]+")

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading AutosInfracaoICMBIO CSV file: %s", err)
			return
		}

		cleanCPF := re.ReplaceAllString(record[9], "")

		autoInfracao := models.AutosInfracaoICMBIO{
			UUID:             uuid.NewString(),
			ID:               strings.Trim(record[0], "\""),
			NumeroAI:         strings.Trim(record[1], "\""),
			Serie:            strings.Trim(record[2], "\""),
			Origem:           strings.Trim(record[3], "\""),
			Tipo:             strings.Trim(record[4], "\""),
			ValorMulta:       strings.Trim(record[5], "\""),
			Embargo:          strings.Trim(record[6], "\""),
			Apreensao:        strings.Trim(record[7], "\""),
			Autuado:          strings.Trim(record[8], "\""),
			CPFCNPJ:          cleanCPF,
			DescricaoAI:      strings.Trim(record[10], "\""),
			DescricaoSancoes: strings.Trim(record[11], "\""),
			Data:             strings.Trim(record[12], "\""),
			Ano:              strings.Trim(record[13], "\""),
			Artigo1:          strings.Trim(record[14], "\""),
			Artigo2:          strings.Trim(record[15], "\""),
			TipoInfracao:     strings.Trim(record[16], "\""),
			NomeUC:           strings.Trim(record[17], "\""),
			CNUC:             strings.Trim(record[18], "\""),
			Municipio:        strings.Trim(record[19], "\""),
			UF:               strings.Trim(record[20], "\""),
			TermosEmbargo:    strings.Trim(record[21], "\""),
			TermosApreensao:  strings.Trim(record[22], "\""),
			Processo:         strings.Trim(record[24], "\""),
			Julgamento:       strings.Trim(record[25], "\""),
		}

		log.Println("Loading to memory AutosInfracaoICMBIO ", autoInfracao.ID)
		autosInfracaoICMBIOData[cleanCPF] = append(autosInfracaoICMBIOData[cleanCPF], autoInfracao)
	}
}

func LoadCSVToMemory_TrabalhoEscravo(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read TrabalhoEscravo input file: %s", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'

	// Expressão regular para remover todos os caracteres não numéricos
	re := regexp.MustCompile("[^0-9]+")

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading TrabalhoEscravo CSV file: %s", err)
			return
		}

		cleanCPF := re.ReplaceAllString(record[4], "")

		trabalhoEscravo := models.TrabalhoEscravo{
			UUID:                         uuid.NewString(),
			ID:                           strings.Trim(record[0], "\""),
			AnoAcaoFiscal:                strings.Trim(record[1], "\""),
			UF:                           strings.Trim(record[2], "\""),
			Empregador:                   strings.Trim(record[3], "\""),
			CNPJCPF:                      cleanCPF,
			Estabelecimento:              strings.Trim(record[5], "\""),
			TrabalhadoresEnvolvidos:      strings.Trim(record[6], "\""),
			CNAE:                         strings.Trim(record[7], "\""),
			DecisaoAdministrativa:        strings.Trim(record[8], "\""),
			InclusaoCadastroEmpregadores: strings.Trim(record[9], "\""),
		}

		log.Println("Loading to memory TrabalhoEscravo ", trabalhoEscravo.ID)
		trabalhoEscravoData[cleanCPF] = append(trabalhoEscravoData[cleanCPF], trabalhoEscravo)
	}
}

func LoadCSVToMemory_Suspensaobama(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read Suspensaobama input file: %s", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'

	// Expressão regular para remover todos os caracteres não numéricos
	re := regexp.MustCompile("[^0-9]+")

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading Suspensaobama CSV file: %s", err)
			return
		}

		cleanCPF := re.ReplaceAllString(record[9], "")

		suspensaobama := models.Suspensaobama{
			UUID:                      uuid.NewString(),
			SEQ_TAD:                   strings.Trim(record[0], "\""),
			STATUS_FORMULARIO:         strings.Trim(record[1], "\""),
			SIT_CANCELADO:             strings.Trim(record[2], "\""),
			NUM_TAD:                   strings.Trim(record[3], "\""),
			SER_TAD:                   strings.Trim(record[4], "\""),
			DAT_TAD:                   strings.Trim(record[5], "\""),
			DAT_IMPRESSAO:             strings.Trim(record[6], "\""),
			NUM_PESSOA_SUSPENSAO:      strings.Trim(record[7], "\""),
			NOM_PESSOA_SUSPENSAO:      strings.Trim(record[8], "\""),
			CPF_CNPJ_PESSOA_SUSPENSAO: cleanCPF,
			NUM_PROCESSO:              strings.Trim(record[10], "\""),
			DES_TAD:                   strings.Trim(record[11], "\""),
			NOM_MUNICIPIO:             strings.Trim(record[13], "\""),
			SIG_UF:                    strings.Trim(record[14], "\""),
			DES_LOCALIZACAO:           strings.Trim(record[15], "\""),
			DES_JUSTIFICATIVA:         strings.Trim(record[18], "\""),
			UNID_CONTROLE:             strings.Trim(record[21], "\""),
			SEQ_AUTO_INFRACAO:         strings.Trim(record[22], "\""),
		}

		log.Println("Loading to memory Suspensaobama ", suspensaobama.SEQ_TAD)
		suspensaobamaData[cleanCPF] = append(suspensaobamaData[cleanCPF], suspensaobama)
	}
}

func LoadCSVToMemory_ApreensaoIbama(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read ApreensaoIbama input file: %s", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'

	// Expressão regular para remover todos os caracteres não numéricos
	re := regexp.MustCompile("[^0-9]+")

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading ApreensaoIbama CSV file: %s", err)
			return
		}

		cleanCPF := re.ReplaceAllString(record[9], "")

		apreensaoibama := models.ApreensaoIbama{
			UUID:                      uuid.NewString(),
			SEQ_TAD:                   record[0],
			STATUS_FORMULARIO:         record[1],
			SIT_CANCELADO:             record[2],
			NUM_TAD:                   record[3],
			SER_TAD:                   record[4],
			DAT_TAD:                   record[5],
			DAT_IMPRESSAO:             record[6],
			NUM_PESSOA_SUSPENSAO:      record[7],
			NOM_PESSOA_SUSPENSAO:      record[8],
			CPF_CNPJ_PESSOA_SUSPENSAO: cleanCPF,
			NUM_PROCESSO:              record[10],
			DES_TAD:                   record[11],
			NOM_MUNICIPIO:             record[13],
			SIG_UF:                    record[14],
			DES_LOCALIZACAO:           record[15],
			SEQ_AUTO_INFRACAO:         record[21],
			SEQ_NOTIFICACAO:           record[22],
		}

		log.Println("Loading to memory ApreensaoIbama ", apreensaoibama.SEQ_TAD)
		apreensaoIbamaData[cleanCPF] = append(apreensaoIbamaData[cleanCPF], apreensaoibama)
	}
}

func LoadAllCSVs(dirPath string) {

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
				case strings.Contains(strings.ToLower(file.Name()), "ceis"):
					LoadCSVToMemory_CEIS(filePath)
				case strings.Contains(strings.ToLower(file.Name()), "cnep"):
					LoadCSVToMemory_CNEP(filePath)
				case strings.Contains(strings.ToLower(file.Name()), "infracao_ibama"):
					LoadCSVToMemory_AutosInfracaoIbama(filePath)
				case strings.Contains(strings.ToLower(file.Name()), "icmbio"):
					LoadCSVToMemory_AutosInfracaoICMBIO(filePath)
				case strings.Contains(strings.ToLower(file.Name()), "escravo"):
					LoadCSVToMemory_TrabalhoEscravo(filePath)
				case strings.Contains(strings.ToLower(file.Name()), "suspensaibama"):
					LoadCSVToMemory_Suspensaobama(filePath)
				case strings.Contains(strings.ToLower(file.Name()), "apreensao"):
					LoadCSVToMemory_ApreensaoIbama(filePath)
				default:
					log.Printf("No suitable function for the file %s", filePath)
				}
			}(file)
		}
	}
	wg.Wait()
}

func LoadCSVToMemory_CB_JSON(filePath string, outputFilePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read input file: %s", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatalf("Unable to create output file: %s", err)
	}
	defer outputFile.Close()

	// add the opening bracket at the beginning of the file
	_, err = outputFile.WriteString("[")
	if err != nil {
		log.Fatalf("Error writing to output file: %s", err)
	}

	isFirst := true
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading CSV file: %s", err)
		}

		tipoPessoa := record[0]
		cpfCnpj := record[1]

		var pep []models.PEP
		var ceis []models.CEIS
		var cnep []models.CNEP
		var aiibama []models.AutosInfracaoIbama
		var aiicmbio []models.AutosInfracaoICMBIO
		var te []models.TrabalhoEscravo
		var sb []models.Suspensaobama
		var ai []models.ApreensaoIbama

		var wg sync.WaitGroup

		wg.Add(8)

		go func() {
			defer wg.Done()
			if tipoPessoa == "PF" {
				re := regexp.MustCompile("[0-9]+")
				cpfCnpjAux := "***." + string(re.Find([]byte(record[1]))) + "-**"
				pep = pepData[cpfCnpjAux]
			}
		}()

		go func() {
			defer wg.Done()
			ceis = ceisData[cpfCnpj]
		}()

		go func() {
			defer wg.Done()
			cnep = cnepData[cpfCnpj]
		}()

		go func() {
			defer wg.Done()
			aiibama = autosInfracaoIbamaData[cpfCnpj]
		}()

		go func() {
			defer wg.Done()
			aiicmbio = autosInfracaoICMBIOData[cpfCnpj]
		}()

		go func() {
			defer wg.Done()
			te = trabalhoEscravoData[cpfCnpj]
		}()

		go func() {
			defer wg.Done()
			sb = suspensaobamaData[cpfCnpj]
		}()

		go func() {
			defer wg.Done()
			ai = apreensaoIbamaData[cpfCnpj]
		}()

		wg.Wait()

		// checks if all objects are nil or not
		if pep == nil && ceis == nil && cnep == nil && aiibama == nil && aiicmbio == nil && te == nil && sb == nil && ai == nil {
			continue
		}

		result := map[string]interface{}{
			"TipoPessoa":            tipoPessoa,
			"CPFCNPJ":               cpfCnpj,
			"PEP":                   pep,
			"CEIS":                  ceis,
			"CNEP":                  cnep,
			"AUTOS_INFRACAO_IBAMA":  aiibama,
			"AUTOS_INFRACAO_ICMBIO": aiicmbio,
			"TRABALHO_ESCRAVO":      te,
			"SUSPENSAO_IBAMA":       sb,
			"APREENSAO_IBAMA":       ai,
		}

		jsonResult, err := json.Marshal(result)
		if err != nil {
			log.Fatalf("Error marshaling JSON: %s", err)
		}

		// add a comma before every record except the first one
		if !isFirst {
			_, err = outputFile.WriteString(",")
			if err != nil {
				log.Fatalf("Error writing to output file: %s", err)
			}
		} else {
			isFirst = false
		}

		_, err = outputFile.Write(jsonResult)
		if err != nil {
			log.Fatalf("Error writing to output file: %s", err)
		}

		_, err = outputFile.WriteString("\n")
		if err != nil {
			log.Fatalf("Error writing to output file: %s", err)
		}
	}

	// add the closing bracket at the end of the file
	_, err = outputFile.WriteString("]")
	if err != nil {
		log.Fatalf("Error writing to output file: %s", err)
	}
}

//func LoadCSVToMemory_CB_JSON(filePath string, outputFilePath string) {
//	file, err := os.Open(filePath)
//	if err != nil {
//		log.Fatalf("Unable to read input file: %s", err)
//	}
//	defer file.Close()
//
//	csvReader := csv.NewReader(file)
//	csvReader.Comma = ';'
//
//	outputFile, err := os.Create(outputFilePath)
//	if err != nil {
//		log.Fatalf("Unable to create output file: %s", err)
//	}
//	defer outputFile.Close()
//
//	for {
//		record, err := csvReader.Read()
//		if err == io.EOF {
//			break
//		}
//		if err != nil {
//			log.Fatalf("Error reading CSV file: %s", err)
//		}
//
//		tipoPessoa := record[0]
//		cpfCnpj := record[1]
//
//		var pep []models.PEP
//		var ceis []models.CEIS
//		var cnep []models.CNEP
//		var aiibama []models.AutosInfracaoIbama
//		var aiicmbio []models.AutosInfracaoICMBIO
//		var te []models.TrabalhoEscravo
//		var sb []models.Suspensaobama
//		var ai []models.ApreensaoIbama
//
//		var wg sync.WaitGroup
//
//		wg.Add(8)
//
//		go func() {
//			defer wg.Done()
//			if tipoPessoa == "PF" {
//				re := regexp.MustCompile("[0-9]+")
//				cpfCnpj := "***." + string(re.Find([]byte(record[1]))) + "-**"
//				pep = pepData[cpfCnpj]
//			}
//		}()
//
//		go func() {
//			defer wg.Done()
//			ceis = ceisData[cpfCnpj]
//		}()
//
//		go func() {
//			defer wg.Done()
//			cnep = cnepData[cpfCnpj]
//		}()
//
//		go func() {
//			defer wg.Done()
//			aiibama = autosInfracaoIbamaData[cpfCnpj]
//		}()
//
//		go func() {
//			defer wg.Done()
//			aiicmbio = autosInfracaoICMBIOData[cpfCnpj]
//		}()
//
//		go func() {
//			defer wg.Done()
//			te = trabalhoEscravoData[cpfCnpj]
//		}()
//
//		go func() {
//			defer wg.Done()
//			sb = suspensaobamaData[cpfCnpj]
//		}()
//
//		go func() {
//			defer wg.Done()
//			ai = apreensaoIbamaData[cpfCnpj]
//		}()
//
//		wg.Wait()
//
//		result := map[string]interface{}{
//			"TipoPessoa":            tipoPessoa,
//			"CPFCNPJ":               cpfCnpj,
//			"PEP":                   pep,
//			"CEIS":                  ceis,
//			"CNEP":                  cnep,
//			"AUTOS_INFRACAO_IBAMA":  aiibama,
//			"AUTOS_INFRACAO_ICMBIO": aiicmbio,
//			"TRABALHO_ESCRAVO":      te,
//			"SUSPENSAO_IBAMA":       sb,
//			"APREENSAO_IBAMA":       ai,
//		}
//
//		jsonResult, err := json.Marshal(result)
//		if err != nil {
//			log.Fatalf("Error marshaling JSON: %s", err)
//		}
//
//		_, err = outputFile.Write(jsonResult)
//		if err != nil {
//			log.Fatalf("Error writing to output file: %s", err)
//		}
//
//		_, err = outputFile.WriteString("\n")
//		if err != nil {
//			log.Fatalf("Error writing to output file: %s", err)
//		}
//	}
//}
