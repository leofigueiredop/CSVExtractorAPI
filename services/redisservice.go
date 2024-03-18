package services

import (
	"CSVExtractor/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/go-redis/redis/v8"
	"log"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})
var client *redisearch.Client

func init() {
	sc := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextField("CPFCNPJ"))
	client = redisearch.NewClient("localhost:6379", "cpfCnpj")

	err := client.CreateIndex(sc)
	if err != nil {
		log.Printf("Error creating index: %v", err)
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

	if err := client.IndexOptions(
		redisearch.DefaultIndexingOptions, doc); err != nil {
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

	if err := client.IndexOptions(
		redisearch.DefaultIndexingOptions, doc); err != nil {
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

func SearchPEPInRedis(key string) (*models.PEP, error) {
	result, err := searchInRedis(key, "PEP")
	if err != nil {
		return nil, err
	}

	model, ok := result.(*models.PEP)
	if !ok {
		return nil, fmt.Errorf("Failed to convert result to model.PEP")
	}

	return model, nil
}

func SearchCEISInRedis(key string) (*models.CEIS, error) {
	result, err := searchInRedis(key, "CEIS")
	if err != nil {
		return nil, err
	}

	model, ok := result.(*models.CEIS)
	if !ok {
		return nil, fmt.Errorf("Failed to convert result to model.CEIS")
	}

	return model, nil
}

func SearchCNEPInRedis(key string) (*models.CNEP, error) {
	result, err := searchInRedis(key, "CNEP")
	if err != nil {
		return nil, err
	}

	model, ok := result.(*models.CNEP)
	if !ok {
		return nil, fmt.Errorf("Failed to convert result to model.CNEP")
	}

	return model, nil
}

func SearchAutosInfracaoIbamaInRedis(key string) (*models.AutosInfracaoIbama, error) {
	result, err := searchInRedis(key, "AutosInfracaoIbama")
	if err != nil {
		return nil, err
	}

	model, ok := result.(*models.AutosInfracaoIbama)
	if !ok {
		return nil, fmt.Errorf("Failed to convert result to model.AutosInfracaoIbama")
	}

	return model, nil
}

func SearchAutosInfracaoICMBIOInRedis(key string) (*models.AutosInfracaoICMBIO, error) {
	result, err := searchInRedis(key, "AutosInfracaoICMBIO")
	if err != nil {
		return nil, err
	}

	model, ok := result.(*models.AutosInfracaoICMBIO)
	if !ok {
		return nil, fmt.Errorf("Failed to convert result to model.AutosInfracaoICMBIO")
	}

	return model, nil
}

func SearchTrabalhoEscravoInRedis(key string) (*models.TrabalhoEscravo, error) {
	result, err := searchInRedis(key, "TrabalhoEscravo")
	if err != nil {
		return nil, err
	}

	model, ok := result.(*models.TrabalhoEscravo)
	if !ok {
		return nil, fmt.Errorf("Failed to convert result to model.TrabalhoEscravo")
	}

	return model, nil
}

func SearchSuspensaobamaInRedis(key string) (*models.Suspensaobama, error) {
	result, err := searchInRedis(key, "Suspensaobama")
	if err != nil {
		return nil, err
	}

	model, ok := result.(*models.Suspensaobama)
	if !ok {
		return nil, fmt.Errorf("Failed to convert result to model.Suspensaobama")
	}

	return model, nil
}

func SearchApreensaoIbamaInRedis(key string) (*models.ApreensaoIbama, error) {
	result, err := searchInRedis(key, "ApreensaoIbama")
	if err != nil {
		return nil, err
	}

	model, ok := result.(*models.ApreensaoIbama)
	if !ok {
		return nil, fmt.Errorf("Failed to convert result to model.ApreensaoIbama")
	}

	return model, nil
}

func searchInRedis(key string, prefix string) (interface{}, error) {
	redisKey := prefix + key
	val, err := rdb.Get(ctx, redisKey).Result()
	if errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("key %s does not exist", redisKey)
	} else if err != nil {
		return nil, fmt.Errorf("Error accessing Redis: %v", err)
	}

	var model interface{}
	switch prefix {
	case "PEP":
		model = &models.PEP{}
	case "CEIS":
		model = &models.CEIS{}
	case "CNEP":
		model = &models.CNEP{}
	case "AutosInfracaoIbama":
		model = &models.AutosInfracaoIbama{}
	case "AutosInfracaoICMBIO":
		model = &models.AutosInfracaoICMBIO{}
	case "TrabalhoEscravo":
		model = &models.TrabalhoEscravo{}
	case "Suspensaobama":
		model = &models.Suspensaobama{}
	case "ApreensaoIbama":
		model = &models.ApreensaoIbama{}
	default:
		return nil, fmt.Errorf("prefix %s not supported", prefix)
	}

	err = json.Unmarshal([]byte(val), model)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling data: %v", err)
	}
	return model, nil
}
