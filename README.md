# Ethereum Transaction Parser

## Overview
This project is a backend service implemented in Go that parses Ethereum blockchain data and provides an interface to interact with transactions for subscribed addresses. It uses the Ethereum JSON-RPC API to fetch blockchain data and allows external users to interact with the parser via HTTP endpoints.

### Key Features:
- Subscribe to Ethereum addresses to track incoming and outgoing transactions.
- Retrieve transactions for a given address.
- Expose a public HTTP API to interact with the parser.

### Caveats:
I am still quite unfamiliar with the Ethereum ecosystem and the Ethereum JSONRPC API so there are quite likely many misunderstandings on how it works thus my implementation here may be incorrect.

## Project Structure
```
ethrpc-parser/
├── api/
│   └── http.go              # HTTP handlers for interacting with the parser
├── entities/
│   └── transaction.go       # Data models for transactions
├── parser/
│   └── parser.go            # Core parser logic to interact with blockchain data
├── rpc/
│   └── eth_rpc.go           # Functions to interact with Ethereum JSON-RPC
├── storage/
│   └── memory.go            # In-memory storage for tracking subscribed addresses and transactions
├── tests/
│   └── (test files)         # Unit tests for core functionality
└── main.go                  # Entry point of the application
```

## Getting Started

### Prerequisites
- [Go](https://golang.org/doc/install) (version 1.18 or later)

### Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/hanifn/ethrpc_parser.git
   cd ethrpc_parser
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```

### Running the Application
You can run the application with the following command:
```sh
go run main.go
```

By default, the service will be accessible at `http://localhost:8080`.

### Endpoints

#### 1. Get Current Block
- **Endpoint**: `/currentBlock`
- **Method**: `GET`
- **Description**: Returns the latest Ethereum block number.
- **Example**:
  ```sh
  curl http://localhost:8080/currentBlock
  ```

#### 2. Subscribe to an Address
- **Endpoint**: `/subscribe`
- **Method**: `GET`
- **Query Parameters**:
  - `address` (required): The Ethereum address to subscribe.
- **Example**:
  ```sh
  curl http://localhost:8080/subscribe?address=0xYourEthereumAddress
  ```

#### 3. Get Transactions for an Address
- **Endpoint**: `/transactions`
- **Method**: `GET`
- **Query Parameters**:
  - `address` (required): The Ethereum address for which to retrieve transactions.
- **Example**:
  ```sh
  curl http://localhost:8080/transactions?address=0xYourEthereumAddress
  ```

## Running Tests
This project includes unit tests for core components. To run all tests, execute:
```sh
go test ./...
```
This will run all tests in the project and provide you with a summary of their status.

