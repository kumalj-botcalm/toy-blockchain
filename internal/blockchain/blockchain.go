package blockchain

import (
	"fmt"
	"strings"
		
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

// AddTransaction adds a transaction to the pending pool.
func (bc *Blockchain) AddTransaction(tx transaction.Transaction) {
	bc.PendingTransactions = append(bc.PendingTransactions, tx)
}

// MinePendingTransactions mines all pending transactions into a new block.
func (bc *Blockchain) MinePendingTransactions() error {

	if len(bc.PendingTransactions) == 0 {
		return fmt.Errorf("no pending transactions")
	}

	lastBlock := bc.Blocks[len(bc.Blocks)-1]

	block := NewBlock(
		len(bc.Blocks),
		bc.PendingTransactions,
		lastBlock.Hash,
	)

	target := strings.Repeat("0", bc.Difficulty)

	for {

		hash, err := block.CalculateHash()
		if err != nil {
			return err
		}

		if strings.HasPrefix(hash, target) {

			block.Hash = hash
			break
		}

		block.Nonce++
	}

	bc.Blocks = append(bc.Blocks, *block)

	bc.PendingTransactions = nil

	return nil
}

