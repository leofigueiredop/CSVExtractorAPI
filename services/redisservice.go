package services

import (
	"CSVExtractor/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func AddToRedis_PEP(pep *models.PEP) {
	jsonPEP, err := json.Marshal(pep)
	if err != nil {
		log.Printf("Unable to marshal object: %v", err)
		return
	}

	err = rdb.Set(ctx, "PEP:"+pep.CPF, jsonPEP, 0).Err()
	if err != nil {
		log.Printf("Unable to save object to Redis: %v", err)
	}
}

func AddToRedis_CEIS(ceis *models.CEIS) {
	jsonCEIS, err := json.Marshal(ceis)
	if err != nil {
		log.Printf("Unable to marshal object: %v", err)
		return
	}

	err = rdb.Set(ctx, "CEIS:"+ceis.CPFCNPJSancionado, jsonCEIS, 0).Err()
	if err != nil {
		log.Printf("Unable to save object to Redis: %v", err)
	}
}

func AddToRedis_CNEP(cnep *models.CNEP) {
	jsonCNEP, err := json.Marshal(cnep)
	if err != nil {
		log.Printf("Unable to marshal object: %v", err)
		return
	}

	err = rdb.Set(ctx, "CNEP:"+cnep.CPFCNPJSancionado, jsonCNEP, 0).Err()
	if err != nil {
		log.Printf("Unable to save object to Redis: %v", err)
	}
}

func AddToRedis_AutosInfracaoIbama(autoInfracao *models.AutosInfracaoIbama) {
	jsonAutoInfracao, err := json.Marshal(autoInfracao)
	if err != nil {
		log.Printf("Unable to marshal object: %v", err)
		return
	}

	err = rdb.Set(ctx, "AutosInfracaoIbama:"+autoInfracao.SeqAutoInfracao, jsonAutoInfracao, 0).Err()
	if err != nil {
		log.Printf("Unable to save object to Redis: %v", err)
	}
}

func AddToRedis_AutosInfracaoICMBIO(autoInfracao *models.AutosInfracaoICMBIO) {
	jsonAutoInfracao, err := json.Marshal(autoInfracao)
	if err != nil {
		log.Printf("Unable to marshal object: %v", err)
		return
	}

	err = rdb.Set(ctx, "AutosInfracaoICMBIO:"+autoInfracao.ID, jsonAutoInfracao, 0).Err()
	if err != nil {
		log.Printf("Unable to save object to Redis: %v", err)
	}
}

func AddToRedis_TrabalhoEscravo(trabalhoEscravo *models.TrabalhoEscravo) {
	jsonTrabalhoEscravo, err := json.Marshal(trabalhoEscravo)
	if err != nil {
		log.Printf("Unable to marshal object: %v", err)
		return
	}

	err = rdb.Set(ctx, "TrabalhoEscravo:"+trabalhoEscravo.ID, jsonTrabalhoEscravo, 0).Err()
	if err != nil {
		log.Printf("Unable to save object to Redis: %v", err)
	}
}

func AddToRedis_Suspensaobama(suspensaobama *models.Suspensaobama) {
	jsonSuspensaobama, err := json.Marshal(suspensaobama)
	if err != nil {
		log.Printf("Unable to marshal object: %v", err)
		return
	}

	err = rdb.Set(ctx, "Suspensaobama:"+suspensaobama.SEQ_TAD, jsonSuspensaobama, 0).Err()
	if err != nil {
		log.Printf("Unable to save object to Redis: %v", err)
	}
}

func AddToRedis_ApreensaoIbama(apreensaoibama *models.ApreensaoIbama) {
	jsonApreensao, err := json.Marshal(apreensaoibama)
	if err != nil {
		log.Printf("Unable to marshal object: %v", err)
		return
	}

	err = rdb.Set(ctx, "ApreensaoIbama:"+apreensaoibama.SEQ_TAD, jsonApreensao, 0).Err()
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
	if err == redis.Nil {
		return nil, fmt.Errorf("Key %s does not exist", redisKey)
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
