package parser

import (
	"encoding/json"
	"ethrpc_parser/storage"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParser_Subsribe(t *testing.T) {
	store := storage.NewMemoryStorage()
	parser := NewParser(store, "testurl")

	success := parser.Subscribe("0x1234567890abcdef")
	if !success {
		t.Errorf("Expected Subscribe to return true, but got false")
	}
}

func TestParser_GetCurrentBlock(t *testing.T) {
	// Set up a mock server to simulate the Ethereum RPC response
	handler := func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      1,
			"result":  "0x4b7", // Mock response block number (1207 in decimal)
		}
		jsonResp, _ := json.Marshal(response)
		w.Write(jsonResp)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	parser := NewParser(storage.NewMemoryStorage(), server.URL)

	// Call GetCurrentBlock and check the result
	blockNumber := parser.GetCurrentBlock()
	if blockNumber != 1207 {
		t.Errorf("Expected block number 1207, got %d", blockNumber)
	}
}

func TestParser_GetTransactions(t *testing.T) {
	// Set up a mock HTTP server to simulate Ethereum RPC responses
	handler := func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      1,
			"result": map[string]interface{}{
				"transactions": []map[string]string{
					{"from": "0x123", "to": "0x456", "value": "1000"},
					{"from": "0x789", "to": "0x456", "value": "500"},
				},
			},
		}
		jsonResp, _ := json.Marshal(response)
		w.Write(jsonResp)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	// Set the RPC URL to the mock server's URL
	rpcUrl := server.URL // Assuming EthRPCURL is used to interact with Ethereum JSON-RPC.

	parser := NewParser(storage.NewMemoryStorage(), rpcUrl) // Create an instance of your parser.

	transactions := parser.GetTransactions("0x456")
	if len(transactions) != 2 {
		t.Errorf("Expected 2 transactions, got %d", len(transactions))
	}
	if transactions[0].Value != "1000" {
		t.Errorf("Expected transaction value '1000', got '%s'", transactions[0].Value)
	}
}
