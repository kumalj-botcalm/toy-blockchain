package blockchain

import (
	"github.com/kumalj-botcalm/toy-blockchain/internal/transaction"
)

// Blockchain represents the blockchain.
type Blockchain struct {
	Blocks              []Block
	PendingTransactions []transaction.Transaction
	Difficulty          int
}

// New creates a new blockchain with a genesis block.
func New(difficulty int) (*Blockchain, error) {

	genesis, err := NewGenesisBlock()
	if err != nil {
		return nil, err
	}

	return &Blockchain{
		Blocks:              []Block{*genesis},
		PendingTransactions: []transaction.Transaction{},
		Difficulty:          difficulty,
	}, nil
}