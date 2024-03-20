package services

import (
	"CSVExtractor/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/go-redis/redis/v8"
	"log"
	"strings"
	"time"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:         "localhost:6380",
	Password:     "",               // no password set
	DB:           0,                // use default DB
	DialTimeout:  60 * time.Second, // tempo de inatividade da configuração da conexão
	ReadTimeout:  90 * time.Second, // tempo de inatividade da leitura
	WriteTimeout: 90 * time.Second, // tempo de inatividade da gravação
})
var client *redisearch.Client

func init() {

	//client = redisearch.NewClient("localhost:6380", "cpfCnpj")
	//
	//sc := redisearch.NewSchema(redisearch.DefaultOptions).
	//	AddField(redisearch.NewTextField("CPFCNPJ"))
	//
	//client.Drop()
	//
	//err := client.CreateIndex(sc)
	//if err != nil {
	//	log.Printf("Error creating index: %v", err)
	//}
}

func SaveRedis() {
	// SAVE
	_, err := rdb.Save(ctx).Result()
	if err != nil {
		log.Printf("Error creating redis file: %v", err)
	}
}

func indexPessoa(pessoa *models.CadastroBasico) {
	doc := redisearch.NewDocument(pessoa.CPF_CNPJ, 1.0)
	doc.Set("CPFCNPJ", pessoa.CPF_CNPJ)

	if err := client.IndexOptions(redisearch.DefaultIndexingOptions, doc); err != nil {
		log.Printf("error indexing document: %v", err)
	}

	jsonPessoa, err := json.Marshal(pessoa)
	if err != nil {
		log.Printf("Unable to marshal object: %v", err)
		return
	}

	err = rdb.Set(ctx, "Pessoa:"+pessoa.CPF_CNPJ, jsonPessoa, 0).Err()
	if err != nil {
		log.Printf("Unable to save object to Redis: %v", err)
	}
}

func indexPEP(pep *models.PEP) {
	doc := redisearch.NewDocument(pep.UUID, 1.0)
	doc.Set("CPFCNPJ", pep.CPF)

	if err := client.IndexOptions(
		redisearch.DefaultIndexingOptions, doc); err != nil {
		log.Printf("error indexing document: %v", err)
	}

	jsonPEP, err := json.Marshal(pep)
	if err != nil {
		log.Printf("Unable to marshal object: %v", err)
		return
	}

	err = rdb.Set(ctx, "PEP:"+pep.UUID, jsonPEP, 0).Err()
	if err != nil {
		log.Printf("Unable to save object to Redis: %v", err)
	}
}

func indexCEIS(ceis *models.CEIS) {

	doc := redisearch.NewDocument(ceis.UUID, 1.0)
	doc.Set("CPFCNPJ", ceis.CPFCNPJSancionado)

	if err := client.Index(doc); err != nil {
		log.Printf("error indexing document: %v", err)
	}

	jsonCEIS, err := json.Marshal(ceis)
	if err != nil {
		log.Printf("Unable to marshal object: %v", err)
		return
	}

	err = rdb.Set(ctx, "CEIS:"+ceis.UUID, jsonCEIS, 0).Err()
	if err != nil {
		log.Printf("Unable to save object to Redis: %v", err)
	}
}

func indexCNEP(cnep *models.CNEP) {

	doc := redisearch.NewDocument(cnep.UUID, 1.0)
	doc.Set("CPFCNPJ", cnep.CPFCNPJSancionado)

	if err := client.Index(doc); err != nil {
		log.Printf("error indexing document: %v", err)
	}

	jsonCNEP, err := json.Marshal(cnep)
	if err != nil {
		log.Printf("Unable to marshal object: %v", err)
		return
	}

	err = rdb.Set(ctx, "CNEP:"+cnep.UUID, jsonCNEP, 0).Err()
	if err != nil {
		log.Printf("Unable to save object to Redis: %v", err)
	}
}

func indexAutoInfracao(autoInfracao *models.AutosInfracaoIbama) {
	doc := redisearch.NewDocument(autoInfracao.UUID, 1.0)
	doc.Set("CPFCNPJ", autoInfracao.CpfCnpjInfrator)

	if err := client.IndexOptions(
		redisearch.DefaultIndexingOptions, doc); err != nil {
		log.Printf("error indexing document: %v", err)
	}

	jsonAutoInfracao, err := json.Marshal(autoInfracao)
	if err != nil {
		log.Printf("Unable to marshal object: %v", err)
		return
	}

	err = rdb.Set(ctx, "AutosInfracaoIbama:"+autoInfracao.UUID, jsonAutoInfracao, 0).Err()
	if err != nil {
		log.Printf("Unable to save object to Redis: %v", err)
	}
}

func indexAutoInfracaoICMBIO(autoInfracao *models.AutosInfracaoICMBIO) {
	doc := redisearch.NewDocument(autoInfracao.UUID, 1.0)
	doc.Set("CPFCNPJ", autoInfracao.CPFCNPJ)

	if err := client.IndexOptions(
		redisearch.DefaultIndexingOptions, doc); err != nil {
		log.Printf("error indexing document: %v", err)
	}

	jsonAutoInfracao, err := json.Marshal(autoInfracao)
	if err != nil {
		log.Printf("Unable to marshal object: %v", err)
		return
	}

	err = rdb.Set(ctx, "AutosInfracaoICMBIO:"+autoInfracao.UUID, jsonAutoInfracao, 0).Err()
	if err != nil {
		log.Printf("Unable to save object to Redis: %v", err)
	}
}

func indexTrabalhoEscravo(trabalhoEscravo *models.TrabalhoEscravo) {
	doc := redisearch.NewDocument(trabalhoEscravo.UUID, 1.0)
	doc.Set("CPFCNPJ", trabalhoEscravo.CNPJCPF)

	if err := client.IndexOptions(
		redisearch.DefaultIndexingOptions, doc); err != nil {
		log.Printf("error indexing document: %v", err)
	}

	jsonTrabalhoEscravo, err := json.Marshal(trabalhoEscravo)
	if err != nil {
		log.Printf("Unable to marshal object: %v", err)
		return
	}
	err = rdb.Set(ctx, "TrabalhoEscravo:"+trabalhoEscravo.UUID, jsonTrabalhoEscravo, 0).Err()
	if err != nil {
		log.Printf("Unable to save object to Redis: %v", err)
	}
}

func indexSuspensaobama(suspensaobama *models.Suspensaobama) {
	doc := redisearch.NewDocument(suspensaobama.UUID, 1.0)
	doc.Set("CPFCNPJ", suspensaobama.CPF_CNPJ_PESSOA_SUSPENSAO)

	if err := client.IndexOptions(
		redisearch.DefaultIndexingOptions, doc); err != nil {
		log.Printf("error indexing document: %v", err)
	}

	jsonSuspensaobama, err := json.Marshal(suspensaobama)
	if err != nil {
		log.Printf("Unable to marshal object: %v", err)
		return
	}

	err = rdb.Set(ctx, "Suspensaobama:"+suspensaobama.UUID, jsonSuspensaobama, 0).Err()
	if err != nil {
		log.Printf("Unable to save object to Redis: %v", err)
	}
}

func indexApreensaoIbama(apreensaoIbama *models.ApreensaoIbama) {
	doc := redisearch.NewDocument(apreensaoIbama.UUID, 1.0)
	doc.Set("CPFCNPJ", apreensaoIbama.CPF_CNPJ_PESSOA_SUSPENSAO)

	if err := client.IndexOptions(
		redisearch.DefaultIndexingOptions, doc); err != nil {
		log.Printf("error indexing document: %v", err)
	}

	jsonApreensao, err := json.Marshal(apreensaoIbama)
	if err != nil {
		log.Printf("Unable to marshal object: %v", err)
		return
	}

	err = rdb.Set(ctx, "ApreensaoIbama:"+apreensaoIbama.UUID, jsonApreensao, 0).Err()
	if err != nil {
		log.Printf("Unable to save object to Redis: %v", err)
	}
}

func SearchCadastroBasicoInRedis(key string) ([]models.Result, error) {
	return searchInRedis(key, "Pessoa")
}

func SearchPEPInRedis(key string) ([]models.Result, error) {
	return searchInRedis(key, "PEP")
}

func SearchCEISInRedis(key string) ([]models.Result, error) {
	return searchInRedis(key, "CEIS")
}

func SearchCNEPInRedis(key string) ([]models.Result, error) {
	return searchInRedis(key, "CNEP")
}

func SearchAutosInfracaoIbamaInRedis(key string) ([]models.Result, error) {
	return searchInRedis(key, "AutosInfracaoIbama")
}

func SearchAutosInfracaoICMBIOInRedis(key string) ([]models.Result, error) {
	return searchInRedis(key, "AutosInfracaoICMBIO")
}

func SearchTrabalhoEscravoInRedis(key string) ([]models.Result, error) {
	return searchInRedis(key, "TrabalhoEscravo")
}

func SearchSuspensaobamaInRedis(key string) ([]models.Result, error) {
	return searchInRedis(key, "Suspensaobama")
}

func SearchApreensaoIbamaInRedis(key string) ([]models.Result, error) {
	return searchInRedis(key, "ApreensaoIbama")
}

func searchInRedis(key string, prefix string) ([]models.Result, error) {
	// search in RediSearch index

	log.Printf(fmt.Sprintf(`@CPFCNPJ:%s*`, key))
	query := redisearch.NewQuery(key)
	docs, _, err := client.Search(query)

	if err != nil {
		return nil, fmt.Errorf("Error searching in RediSearch: %v", err)
	}

	var results []models.Result
	for _, doc := range docs {
		if strings.HasPrefix(doc.Id, prefix) {
			val, err := rdb.Get(ctx, doc.Id).Result()
			if err != nil {
				return nil, fmt.Errorf("Error accessing Redis: %v", err)
			}

			results = append(results, models.Result{
				Key:   doc.Id,
				Index: prefix,
				Data:  json.RawMessage(val),
			})
		}
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("Key %s does not exist", key)
	}

	return results, nil
}
