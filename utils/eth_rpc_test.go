package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCurrentBlockNumber(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      1,
			"result":  "0x4b7", // Mock response block number
		}
		jsonResp, _ := json.Marshal(response)
		w.Write(jsonResp)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	blockNumber, err := GetCurrentBlockNumber(server.URL) // Assuming this function exists in your RPC client
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if blockNumber != 1207 {
		t.Errorf("Expected block number 1207, got %d", blockNumber)
	}
}

func TestGetTransaction(t *testing.T) {
	// Set up a mock server to simulate Ethereum JSON-RPC response
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Mock response with a couple of transactions
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

	// Call GetTransaction with the mock server URL
	transactions, err := GetTransaction(server.URL, "latest") // Replace "latest" with the appropriate parameter

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Validate the results
	if len(transactions) != 2 {
		t.Errorf("Expected 2 transactions, got %d", len(transactions))
	}

	if transactions[0].From != "0x123" || transactions[0].To != "0x456" || transactions[0].Value != "1000" {
		t.Errorf("Unexpected transaction values. Got %+v", transactions[0])
	}

	if transactions[1].From != "0x789" || transactions[1].To != "0x456" || transactions[1].Value != "500" {
		t.Errorf("Unexpected transaction values. Got %+v", transactions[1])
	}
}

func TestGetTransaction_InvalidJSON(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{invalid json}`)) // Return invalid JSON to trigger an error
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	_, err := GetTransaction(server.URL, "latest")
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}
