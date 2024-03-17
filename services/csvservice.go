package services

import (
	"CSVExtractor/models"
	"encoding/csv"
	"github.com/google/uuid"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var lock = &sync.Mutex{}

func LoadCSVToRedis_PEP(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read input file: %s", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading CSV file: %s", err)
		}

		pep := models.PEP{
			UUID:                  uuid.NewString(),
			CPF:                   strings.Trim(record[0], "\""),
			Nome_PEP:              strings.Trim(record[1], "\""),
			Sigla_Funcao:          strings.Trim(record[2], "\""),
			Descricao_Funcao:      strings.Trim(record[3], "\""),
			Nivel_Funcao:          strings.Trim(record[4], "\""),
			Nome_Orgao:            strings.Trim(record[5], "\""),
			Data_Inicio_Exercicio: strings.Trim(record[6], "\""),
			Data_Fim_Exercicio:    strings.Trim(record[7], "\""),
			Data_Fim_Carencia:     strings.Trim(record[8], "\""),
		}

		log.Println("Saving to redis PEP", pep.CPF)
		AddToRedis_PEP(&pep)
	}
}

func LoadCSVToRedis_CEIS(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read input file: %s", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading CSV file: %s", err)
		}

		ceis := models.CEIS{
			UUID:                          uuid.NewString(),
			Cadastro:                      strings.Trim(record[0], "\""),
			CodigoSancao:                  strings.Trim(record[1], "\""),
			TipoPessoa:                    strings.Trim(record[2], "\""),
			CPFCNPJSancionado:             strings.Trim(record[3], "\""),
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
			EsferaOrgaoSancionador:        strings.Trim(record[19], "\""),
			FundamentacaoLegal:            strings.Trim(record[20], "\""),
		}

		log.Println("Saving to redis CEIS", ceis.CPFCNPJSancionado)
		AddToRedis_CEIS(&ceis)
	}
}

func LoadCSVToRedis_CNEP(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read CNEP input file: %s", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading CNEP CSV file: %s", err)
			return
		}

		cnep := models.CNEP{
			UUID:                          uuid.NewString(),
			Cadastro:                      strings.Trim(record[0], "\""),
			CodigoSancao:                  strings.Trim(record[1], "\""),
			TipoPessoa:                    strings.Trim(record[2], "\""),
			CPFCNPJSancionado:             strings.Trim(record[3], "\""),
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
			EsferaOrgaoSancionador:        strings.Trim(record[20], "\""),
			FundamentacaoLegal:            strings.Trim(record[21], "\""),
		}

		log.Println("Saving to redis CNEP", cnep.CPFCNPJSancionado)
		AddToRedis_CNEP(&cnep)
	}
}

func LoadCSVToRedis_AutosInfracaoIbama(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read AutosInfracaoIbama input file: %s", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error reading AutosInfracaoIbama CSV file: %s", err) // changed here
			continue                                                         // added here

		} else {

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
				//FormaEntrega:               strings.Trim(record[11], "\""),
				DatCienciaAutuacao: strings.Trim(record[12], "\""),
				CodMunicipio:       strings.Trim(record[13], "\""),
				Municipio:          strings.Trim(record[14], "\""),
				Uf:                 strings.Trim(record[15], "\""),
				NumProcesso:        strings.Trim(record[16], "\""),
				CodInfracao:        strings.Trim(record[17], "\""),
				DesInfracao:        strings.Trim(record[18], "\""),
				TipoInfracao:       strings.Trim(record[19], "\""),
				NomeInfrator:       strings.Trim(record[20], "\""),
				CpfCnpjInfrator:    strings.Trim(record[21], "\""),
				//QtdArea:                    strings.Trim(record[22], "\""),
				//InfracaoArea:               strings.Trim(record[23], "\""),
				//DesOutrosTipoArea:          strings.Trim(record[24], "\""),
				//ClassificacaoArea:          strings.Trim(record[25], "\""),
				//NumLatitudeAuto:            strings.Trim(record[26], "\""),
				//NumLongitudeAuto:           strings.Trim(record[27], "\""),
				DesLocalInfracao:     strings.Trim(record[28], "\""),
				NotificacaoVinculada: strings.Trim(record[29], "\""),
				AcaoFiscalizatoria:   strings.Trim(record[30], "\""),
				UnidControle:         strings.Trim(record[31], "\""),
				TipoAcao:             strings.Trim(record[32], "\""),
				Operacao:             strings.Trim(record[33], "\""),
				//DenunciaSisliv:             strings.Trim(record[34], "\""),
				//OrdemFiscalizacao:          strings.Trim(record[35], "\""),
				//SolicitacaoRecurso:         strings.Trim(record[36], "\""),
				//OperacaoSolRecurso:         strings.Trim(record[37], "\""),
				DatLancamento: strings.Trim(record[38], "\""),
				//DatUltAlteracao:            strings.Trim(record[39], "\""),
				//TipoUltAlteracao:           strings.Trim(record[40], "\""),
				//UltimaAtualizacaoRelatorio: strings.Trim(record[41], "\""),
			}

			log.Println("Saving to redis AutosInfracaoIbama ", autoInfracao.SeqAutoInfracao)
			AddToRedis_AutosInfracaoIbama(&autoInfracao)
		}
	}
}

func LoadCSVToRedis_AutosInfracaoICMBIO(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read AutosInfracaoICMBIO input file: %s", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading AutosInfracaoICMBIO CSV file: %s", err)
			return
		}

		autoInfracao := models.AutosInfracaoICMBIO{
			UUID:              uuid.NewString(),
			ID:                strings.Trim(record[0], "\""),
			NumeroAI:          strings.Trim(record[1], "\""),
			Serie:             strings.Trim(record[2], "\""),
			Origem:            strings.Trim(record[3], "\""),
			Tipo:              strings.Trim(record[4], "\""),
			ValorMulta:        strings.Trim(record[5], "\""),
			Embargo:           strings.Trim(record[6], "\""),
			Apreensao:         strings.Trim(record[7], "\""),
			Autuado:           strings.Trim(record[8], "\""),
			CPFCNPJ:           strings.Trim(record[9], "\""),
			DescricaoAI:       strings.Trim(record[10], "\""),
			DescricaoSancoes:  strings.Trim(record[11], "\""),
			Data:              strings.Trim(record[12], "\""),
			Ano:               strings.Trim(record[13], "\""),
			Artigo1:           strings.Trim(record[14], "\""),
			Artigo2:           strings.Trim(record[15], "\""),
			TipoInfracao:      strings.Trim(record[16], "\""),
			NomeUC:            strings.Trim(record[17], "\""),
			CNUC:              strings.Trim(record[18], "\""),
			Municipio:         strings.Trim(record[19], "\""),
			UF:                strings.Trim(record[20], "\""),
			TermosEmbargo:     strings.Trim(record[21], "\""),
			TermosApreensao:   strings.Trim(record[22], "\""),
			OrdemFiscalizacao: strings.Trim(record[23], "\""),
			Processo:          strings.Trim(record[24], "\""),
			Julgamento:        strings.Trim(record[25], "\""),
		}

		log.Println("Saving to redis AutosInfracaoICMBIO ", autoInfracao.ID)
		AddToRedis_AutosInfracaoICMBIO(&autoInfracao)
	}
}

func LoadCSVToRedis_TrabalhoEscravo(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read TrabalhoEscravo input file: %s", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading TrabalhoEscravo CSV file: %s", err)
			return
		}

		trabalhoEscravo := models.TrabalhoEscravo{
			UUID:                         uuid.NewString(),
			ID:                           strings.Trim(record[0], "\""),
			AnoAcaoFiscal:                strings.Trim(record[1], "\""),
			UF:                           strings.Trim(record[2], "\""),
			Empregador:                   strings.Trim(record[3], "\""),
			CNPJCPF:                      strings.Trim(record[4], "\""),
			Estabelecimento:              strings.Trim(record[5], "\""),
			TrabalhadoresEnvolvidos:      strings.Trim(record[6], "\""),
			CNAE:                         strings.Trim(record[7], "\""),
			DecisaoAdministrativa:        strings.Trim(record[8], "\""),
			InclusaoCadastroEmpregadores: strings.Trim(record[9], "\""),
		}

		log.Println("Saving to redis TrabalhoEscravo ", trabalhoEscravo.ID)
		AddToRedis_TrabalhoEscravo(&trabalhoEscravo)
	}
}

func LoadCSVToRedis_Suspensaobama(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read Suspensaobama input file: %s", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading Suspensaobama CSV file: %s", err)
			return
		}

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
			CPF_CNPJ_PESSOA_SUSPENSAO: strings.Trim(record[9], "\""),
			NUM_PROCESSO:              strings.Trim(record[10], "\""),
			DES_TAD:                   strings.Trim(record[11], "\""),
			COD_MUNICIPIO:             strings.Trim(record[12], "\""),
			NOM_MUNICIPIO:             strings.Trim(record[13], "\""),
			SIG_UF:                    strings.Trim(record[14], "\""),
			DES_LOCALIZACAO:           strings.Trim(record[15], "\""),
			//NUM_LONGITUDE_TAD:            strings.Trim(record[16], "\""),
			//NUM_LATITUDE_TAD:             strings.Trim(record[17], "\""),
			DES_JUSTIFICATIVA:       strings.Trim(record[18], "\""),
			FORMA_ENTREGA:           strings.Trim(record[19], "\""),
			UNID_APRESENTACAO:       strings.Trim(record[20], "\""),
			UNID_CONTROLE:           strings.Trim(record[21], "\""),
			SEQ_AUTO_INFRACAO:       strings.Trim(record[22], "\""),
			SEQ_NOTIFICACAO:         strings.Trim(record[23], "\""),
			SEQ_ACAO_FISCALIZATORIA: strings.Trim(record[24], "\""),
			SEQ_ORDEM_FISCALIZACAO:  strings.Trim(record[25], "\""),
			NUM_ORDEM_FISCALIZACAO:  strings.Trim(record[26], "\""),
			//SEQ_SOLICITACAO_RECURSO:      strings.Trim(record[27], "\""),
			//NUM_SOLICITACAO_RECURSO:      strings.Trim(record[28], "\""),
			//OPERACAO_SOL_RECURSO:         strings.Trim(record[29], "\""),
			//DAT_ALTERACAO:                strings.Trim(record[30], "\""),
			//TIPO_ALTERACAO:               strings.Trim(record[31], "\""),
			//JUSTIFICATIVA_ALTERACAO:      strings.Trim(record[32], "\""),
			//ULTIMA_ATUALIZACAO_RELATORIO: strings.Trim(record[33], "\""),
		}

		log.Println("Saving to redis Suspensaobama ", suspensaobama.SEQ_TAD)
		AddToRedis_Suspensaobama(&suspensaobama)
	}
}

func LoadCSVToRedis_ApreensaoIbama(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read ApreensaoIbama input file: %s", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading ApreensaoIbama CSV file: %s", err)
			return
		}

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
			CPF_CNPJ_PESSOA_SUSPENSAO: record[9],
			NUM_PROCESSO:              record[10],
			DES_TAD:                   record[11],
			COD_MUNICIPIO:             record[12],
			NOM_MUNICIPIO:             record[13],
			SIG_UF:                    record[14],
			DES_LOCALIZACAO:           record[15],
			//NUM_LONGITUDE_TAD:            record[16],
			//NUM_LATITUDE_TAD:             record[17],
			FORMA_ENTREGA:           record[18],
			UNID_APRESENTACAO:       record[19],
			UNID_CONTROLE:           record[20],
			SEQ_AUTO_INFRACAO:       record[21],
			SEQ_NOTIFICACAO:         record[22],
			SEQ_ACAO_FISCALIZATORIA: record[23],
			SEQ_ORDEM_FISCALIZACAO:  record[24],
			NUM_ORDEM_FISCALIZACAO:  record[25],
			//SEQ_SOLICITACAO_RECURSO:      record[26],
			//NUM_SOLICITACAO_RECURSO:      record[27],
			//OPERACAO_SOL_RECURSO:         record[28],
			//DAT_ALTERACAO:                record[29],
			//TIPO_ALTERACAO:               record[30],
			//JUSTIFICATIVA_ALTERACAO:      record[31],
			//ULTIMA_ATUALIZACAO_RELATORIO: record[32],
		}

		log.Println("Saving to redis ApreensaoIbama ", apreensaoibama.SEQ_TAD)
		AddToRedis_ApreensaoIbama(&apreensaoibama)
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
					LoadCSVToRedis_PEP(filePath)
				case strings.Contains(strings.ToLower(file.Name()), "ceis"):
					LoadCSVToRedis_CEIS(filePath)
				case strings.Contains(strings.ToLower(file.Name()), "cnep"):
					LoadCSVToRedis_CNEP(filePath)
				case strings.Contains(strings.ToLower(file.Name()), "infracao_ibama"):
					LoadCSVToRedis_AutosInfracaoIbama(filePath)
				case strings.Contains(strings.ToLower(file.Name()), "icmbio"):
					LoadCSVToRedis_AutosInfracaoICMBIO(filePath)
				case strings.Contains(strings.ToLower(file.Name()), "escravo"):
					LoadCSVToRedis_TrabalhoEscravo(filePath)
				case strings.Contains(strings.ToLower(file.Name()), "suspensaibama"):
					LoadCSVToRedis_Suspensaobama(filePath)
				case strings.Contains(strings.ToLower(file.Name()), "apreensao"):
					LoadCSVToRedis_ApreensaoIbama(filePath)
				default:
					log.Printf("No suitable function for the file %s", filePath)
				}
			}(file)
		}
	}
	wg.Wait()
}
