package main

import (
	"ethrpc_parser/api"
	"ethrpc_parser/parser"
	"ethrpc_parser/storage"
	"fmt"
	"net/http"
)

const EthRPCURL = "https://ethereum-rpc.publicnode.com"

func main() {
	// init new parser with in-memory storage
	memoryStorage := storage.NewMemoryStorage()
	ethParser := parser.NewParser(memoryStorage, EthRPCURL)

	handlers := &api.API{Parser: ethParser}

	http.HandleFunc("/currentBlock", handlers.GetCurrentBlock)
	http.HandleFunc("/subscribe", handlers.Subscribe)
	http.HandleFunc("/transactions", handlers.GetTransactions)

	// start the http server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
