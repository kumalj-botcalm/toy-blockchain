package blockchain

import "github.com/kumalj-botcalm/toy-blockchain/internal/transaction"

// Blockchain represents the blockchain.
type Blockchain struct {
	Blocks              []Block
	PendingTransactions []transaction.Transaction
	Difficulty          int
}
