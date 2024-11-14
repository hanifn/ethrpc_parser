package parser

type Transaction struct {
	From  string
	To    string
	Value string
}

type Parser interface {
	GetCurrentBlock() int
	Subscribe(address string) bool
	GetTransactions(address string) []Transaction
}
