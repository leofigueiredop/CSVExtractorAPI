package handlers

import (
	"CSVExtractor/services"
	"encoding/json"
	"net/http"
)

// @Summary Search PEP
// @Description Get PEP by key
// @ID get-pep
// @Accept json
// @Produce json
// @Param key query string true "PEP key"
// @Router /search/pep [get]
func SearchHandlerPEP(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key parameter is missing", http.StatusBadRequest)
		return
	}

	model, err := services.SearchPEPInRedis(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(model)
}

// @Summary Search CEIS
// @Description Get CEIS by key
// @ID get-ceis
// @Accept json
// @Produce json
// @Param key query string true "CEIS key"
// @Router /search/ceis [get]
func SearchHandlerCEIS(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key parameter is missing", http.StatusBadRequest)
		return
	}

	model, err := services.SearchCEISInRedis(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(model)
}

// @Summary Search CNEP
// @Description Get CNEP by key
// @ID get-cnep
// @Accept json
// @Produce json
// @Param key query string true "CNEP key"
// @Router /search/cnep [get]
func SearchHandlerCNEP(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key parameter is missing", http.StatusBadRequest)
		return
	}

	model, err := services.SearchCNEPInRedis(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(model)
}

// @Summary Search Autos Infracao Ibama
// @Description Get Autos Infracao Ibama by key
// @ID get-autos-infracao-ibama
// @Accept json
// @Produce json
// @Param key query string true "Autos Infracao Ibama key"
// @Router /search/AutosInfracaoIbama [get]
func SearchHandlerAutosInfracaoIbama(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key parameter is missing", http.StatusBadRequest)
		return
	}

	model, err := services.SearchAutosInfracaoIbamaInRedis(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(model)
}

// @Summary Search Autos Infracao ICMBIO
// @Description Get Autos Infracao ICMBIO by key
// @ID get-autos-infracao-icmbio
// @Accept json
// @Produce json
// @Param key query string true "Autos Infracao ICMBIO key"
// @Router /search/AutosInfracaoICMBIO [get]
func SearchHandlerAutosInfracaoICMBIO(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key parameter is missing", http.StatusBadRequest)
		return
	}

	model, err := services.SearchAutosInfracaoICMBIOInRedis(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(model)
}

// @Summary Search Trabalho Escravo
// @Description Get Trabalho Escravo by key
// @ID get-trabalho-escravo
// @Accept json
// @Produce json
// @Param key query string true "Trabalho Escravo key"
// @Router /search/TrabalhoEscravo [get]
func SearchHandlerTrabalhoEscravo(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key parameter is missing", http.StatusBadRequest)
		return
	}

	model, err := services.SearchTrabalhoEscravoInRedis(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(model)
}

// @Summary Search Suspensao bama
// @Description Get Suspensao bama by key
// @ID get-suspensao-bama
// @Accept json
// @Produce json
// @Param key query string true "Suspensao bama key"
// @Router /search/Suspensaobama [get]
func SearchHandlerSuspensaobama(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key parameter is missing", http.StatusBadRequest)
		return
	}

	model, err := services.SearchSuspensaobamaInRedis(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(model)
}

// @Summary Search Apreensao Ibama
// @Description Get Apreensao Ibama by key
// @ID get-apreensao-ibama
// @Accept json
// @Produce json
// @Param key query string true "Apreensao Ibama key"
// @Router /search/ApreensaoIbama [get]
func SearchHandlerApreensaoIbama(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key parameter is missing", http.StatusBadRequest)
		return
	}

	model, err := services.SearchApreensaoIbamaInRedis(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(model)
}
