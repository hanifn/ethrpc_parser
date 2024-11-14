package storage

import (
	"ethrpc_parser/entities"
	"testing"
)

func TestMemoryStorage_Subscribe(t *testing.T) {
	storage := NewMemoryStorage()

	// Test subscribing to an address
	storage.Subscribe("0x12345")
	if !storage.Addresses["0x12345"] {
		t.Error("Expected address to be subscribed, but it wasn't")
	}
}

func TestMemoryStorage_AddTransaction(t *testing.T) {
	storage := NewMemoryStorage()

	tx := entities.Transaction{From: "0xabc", To: "0xdef", Value: "123"}
	storage.AddTransaction("0x12345", tx)

	transactions := storage.GetTransactions("0x12345")
	if len(transactions) != 1 {
		t.Errorf("Expected 1 transaction, got %d", len(transactions))
	}
	if transactions[0].Value != "123" {
		t.Errorf("Expected transaction value '123', got '%s'", transactions[0].Value)
	}
}
