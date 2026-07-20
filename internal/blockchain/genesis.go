package blockchain

import (
	"github.com/kumalj-botcalm/toy-blockchain/internal/transaction"
)

const (
	GenesisPreviousHash       = "0000000000000000000000000000000000000000000000000000000000000000"
	GenesisTimestamp    int64 = 1721174400
)

// NewGenesisBlock creates the first block in the blockchain.
func NewGenesisBlock() *Block {

	return &Block{
		Index:        0,
		Timestamp:    GenesisTimestamp,
		Transactions: []transaction.Transaction{},
		PreviousHash: GenesisPreviousHash,
		Nonce:        0,
	}
}
