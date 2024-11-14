package storage

import (
	"ethrpc_parser/entities"
	"sync"
)

type MemoryStorage struct {
	Addresses    map[string]bool
	Transactions map[string][]entities.Transaction
	mu           sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		Addresses:    make(map[string]bool),
		Transactions: make(map[string][]entities.Transaction),
	}
}

func (ms *MemoryStorage) Subscribe(address string) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.Addresses[address] = true
}

func (ms *MemoryStorage) AddTransaction(address string, tx entities.Transaction) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.Transactions[address] = append(ms.Transactions[address], tx)
}

func (ms *MemoryStorage) GetTransactions(address string) []entities.Transaction {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	return ms.Transactions[address]
}
