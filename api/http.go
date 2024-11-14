package api

import (
	"encoding/json"
	"ethrpc_parser/parser"
	"net/http"
)

type API struct {
	Parser parser.Parser
}

func (api *API) GetCurrentBlock(w http.ResponseWriter, r *http.Request) {
	block := api.Parser.GetCurrentBlock()
	json.NewEncoder(w).Encode(map[string]int{"currentBlock": block})
}

func (api *API) Subscribe(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Missing address parameter", http.StatusBadRequest)
		return
	}
	api.Parser.Subscribe(address)
	w.WriteHeader(http.StatusOK)
}

func (api *API) GetTransactions(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Missing address parameter", http.StatusBadRequest)
		return
	}
	transactions := api.Parser.GetTransactions(address)
	json.NewEncoder(w).Encode(transactions)
}
