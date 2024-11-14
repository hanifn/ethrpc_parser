package utils

import (
	"bytes"
	"encoding/json"
	"ethrpc_parser/entities"
	"net/http"
	"strconv"
)

const EthRPCURL = "https://ethereum-rpc.publicnode.com"

type ReqBody struct {
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      int           `json:"id"`
}

func newReqBody() ReqBody {
	return ReqBody{
		JsonRpc: "2.0",
		Params:  []interface{}{},
		Id:      0,
	}
}

func GetCurrentBlockNumber() (int, error) {
	body := newReqBody()
	body.Method = "eth_blockNumber"
	reqBody, err := json.Marshal(body)
	if err != nil {
		return 0, err
	}

	resp, err := http.Post(EthRPCURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result struct {
		Result string `json:"result"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return 0, err
	}

	// Convert from hex to int
	blockNumber, err := strconv.ParseInt(result.Result[2:], 16, 64)
	if err != nil {
		return 0, err
	}
	return int(blockNumber), nil
}

func GetTransaction(address string) ([]entities.Transaction, error) {
	body := newReqBody()
	body.Method = "eth_getBlockByNumber"
	body.Params = append(body.Params, address, true)
	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// Get transactions from eth_getBlockByNumber method of Ethereum RPC API
	resp, err := http.Post(EthRPCURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Result struct {
			Transactions []entities.Transaction `json:"transactions"`
		} `json:"result"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Result.Transactions, nil
}
