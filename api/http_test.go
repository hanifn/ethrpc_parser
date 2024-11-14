package api

import (
	"encoding/json"
	"ethrpc_parser/entities"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mock struct {
	getCurrentBlock func() int
	subscribe       func(string) bool
	getTransactions func(string) []entities.Transaction
}

func (m *mock) GetCurrentBlock() int          { return m.getCurrentBlock() }
func (m *mock) Subscribe(address string) bool { return m.subscribe(address) }
func (m *mock) GetTransactions(address string) []entities.Transaction {
	return m.getTransactions(address)
}

func TestAPI_GetCurrentBlock(t *testing.T) {
	// Create a mock parser that returns a fixed value for GetCurrentBlock
	mockParser := &mock{
		getCurrentBlock: func() int {
			return 1207
		},
	}
	api := &API{Parser: mockParser}

	req := httptest.NewRequest(http.MethodGet, "/currentBlock", nil)
	w := httptest.NewRecorder()

	// Call the handler
	api.GetCurrentBlock(w, req)

	// Validate response
	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", res.Status)
	}

	var response map[string]int
	json.NewDecoder(res.Body).Decode(&response)

	if response["currentBlock"] != 1207 {
		t.Errorf("Expected currentBlock 1207, got %d", response["currentBlock"])
	}
}

func TestAPI_GetCurrentBlock_Fail(t *testing.T) {
	// Create a mock parser that returns a fixed value for GetCurrentBlock
	mockParser := &mock{
		getCurrentBlock: func() int {
			return -1
		},
	}
	api := &API{Parser: mockParser}

	req := httptest.NewRequest(http.MethodGet, "/currentBlock", nil)
	w := httptest.NewRecorder()

	// Call the handler
	api.GetCurrentBlock(w, req)

	// Validate response
	res := w.Result()
	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status Internal Server Error, got %v", res.Status)
	}

	expectedError := "Failed to get current block"
	body := strings.TrimSpace(w.Body.String())
	if body != expectedError {
		t.Errorf("Expected error message %q, got %q", expectedError, body)
	}
}

func TestAPI_Subscribe_Success(t *testing.T) {
	// Create a mock parser that always returns true for Subscribe
	mockParser := &mock{
		subscribe: func(address string) bool {
			return true
		},
	}
	api := &API{Parser: mockParser}

	req := httptest.NewRequest(http.MethodGet, "/subscribe?address=0x12345", nil)
	w := httptest.NewRecorder()

	// Call the handler
	api.Subscribe(w, req)

	// Validate response
	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", res.Status)
	}
}

func TestAPI_Subscribe_MissingAddress(t *testing.T) {
	mockParser := &mock{}
	api := &API{Parser: mockParser}

	req := httptest.NewRequest(http.MethodGet, "/subscribe", nil)
	w := httptest.NewRecorder()

	// Call the handler
	api.Subscribe(w, req)

	// Validate response
	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status BadRequest, got %v", res.Status)
	}

	expectedError := "Missing address parameter"
	body := strings.TrimSpace(w.Body.String())
	if body != expectedError {
		t.Errorf("Expected error message %q, got %q", expectedError, body)
	}
}

func TestAPI_GetTransactions_MissingAddress(t *testing.T) {
	mockParser := &mock{}
	api := &API{Parser: mockParser}

	req := httptest.NewRequest(http.MethodGet, "/transactions", nil)
	w := httptest.NewRecorder()

	// Call the handler
	api.GetTransactions(w, req)

	// Validate response
	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status BadRequest, got %v", res.Status)
	}

	expectedError := "Missing address parameter"
	body := strings.TrimSpace(w.Body.String())
	if body != expectedError {
		t.Errorf("Expected error message %q, got %q", expectedError, body)
	}
}

func TestAPI_GetTransactions_Success(t *testing.T) {
	// Mock the parser's GetTransactions method to return a fixed set of transactions
	mockParser := &mock{
		getTransactions: func(address string) []entities.Transaction {
			return []entities.Transaction{
				{From: "0xabc", To: "0xdef", Value: "1000"},
				{From: "0x123", To: "0x456", Value: "500"},
			}
		},
	}
	api := &API{Parser: mockParser}

	req := httptest.NewRequest(http.MethodGet, "/transactions?address=0x123", nil)
	w := httptest.NewRecorder()

	// Call the handler
	api.GetTransactions(w, req)

	// Validate response
	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", res.Status)
	}

	var transactions []entities.Transaction
	err := json.NewDecoder(res.Body).Decode(&transactions)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	// Check if the transactions match the expected values
	if len(transactions) != 2 {
		t.Errorf("Expected 2 transactions, got %d", len(transactions))
	}

	if transactions[0].From != "0xabc" || transactions[0].To != "0xdef" || transactions[0].Value != "1000" {
		t.Errorf("Unexpected transaction values. Got %+v", transactions[0])
	}

	if transactions[1].From != "0x123" || transactions[1].To != "0x456" || transactions[1].Value != "500" {
		t.Errorf("Unexpected transaction values. Got %+v", transactions[1])
	}
}
