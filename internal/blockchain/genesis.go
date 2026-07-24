package blockchain

import (
	"github.com/kumalj-botcalm/toy-blockchain/internal/merkle"
	"github.com/kumalj-botcalm/toy-blockchain/internal/transaction"
)

const (
	GenesisPreviousHash       = "0000000000000000000000000000000000000000000000000000000000000000"
	GenesisTimestamp    int64 = 1721174400
)

// NewGenesisBlock creates the deterministic genesis block.
func NewGenesisBlock() (*Block, error) {

	block := &Block{
		Index:        0,
		Timestamp:    GenesisTimestamp,
		Transactions: []transaction.Transaction{},
		PreviousHash: GenesisPreviousHash,
		Nonce:        0,
	}

	root, err := merkle.BuildMerkleRoot(block.Transactions)
	if err != nil {
		return nil, err
	}

	block.MerkleRoot = root

	hash, err := block.CalculateHash()
	if err != nil {
		return nil, err
	}

	block.Hash = hash

	return block, nil
}
