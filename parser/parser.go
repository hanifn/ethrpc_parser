package parser

import (
	"ethrpc_parser/entities"
	"ethrpc_parser/storage"
	"ethrpc_parser/utils"
	"fmt"
)

type Parser interface {
	// last parsed block
	GetCurrentBlock() int

	// add address to observer
	Subscribe(address string) bool

	// list of inbound or outbound transactions for an address
	GetTransactions(address string) []entities.Transaction
}

type EthParser struct {
	store  *storage.MemoryStorage
	rpcUrl string
}

// NewParser returns a new instance of EthParser
func NewParser(store *storage.MemoryStorage, rpcUrl string) *EthParser {
	return &EthParser{store, rpcUrl}
}

// GetCurrentBlock retrieves the last parsed block from the Ethereum JSONRPC API
func (p *EthParser) GetCurrentBlock() int {
	blockNum, err := utils.GetCurrentBlockNumber(p.rpcUrl)
	if err != nil {
		fmt.Println("Error getting current block number")
		return -1
	}
	return blockNum
}

// Subscribe adds the specified address to the observer
func (p *EthParser) Subscribe(address string) bool {
	p.store.Subscribe(address)
	return true
}

// GetTransactions retrieves a list of inbound and outbound transactions for an address from the Ethereum JSONRPC API
func (p *EthParser) GetTransactions(address string) []entities.Transaction {
	transactions, err := utils.GetTransaction(p.rpcUrl, address)
	if err != nil {
		fmt.Println("Error getting transactions")
		return nil
	}

	// Save all transactions to storage
	for _, tx := range transactions {
		// Save to storage
		p.store.AddTransaction(address, tx)
	}

	return transactions
}
