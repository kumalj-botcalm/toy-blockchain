package blockchain

import (
	"fmt"
	"time"

	"github.com/kumalj-botcalm/toy-blockchain/internal/miner"
	"github.com/kumalj-botcalm/toy-blockchain/internal/transaction"
)

// Blockchain represents the blockchain.
type Blockchain struct {
	Blocks              []Block
	PendingTransactions []transaction.Transaction
	Difficulty          int
	MaxBlockSize        int
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
		MaxBlockSize:        DefaultBlockSize,
	}, nil
}

// AddTransaction adds a transaction to the pending pool.
func (bc *Blockchain) AddTransaction(tx transaction.Transaction) error {

	// SYSTEM account can always create funds.
	if tx.Sender == SystemAccount {
		bc.PendingTransactions = append(bc.PendingTransactions, tx)
		return nil
	}

	balances := bc.AvailableBalances()

	if balances[tx.Sender] < tx.Amount {
		return fmt.Errorf("insufficient balance")
	}

	bc.PendingTransactions = append(bc.PendingTransactions, tx)

	return nil
}

// MinePendingTransactions mines pending transactions into a block.
func (bc *Blockchain) MinePendingTransactions() error {

	if len(bc.PendingTransactions) == 0 {
		return fmt.Errorf("no pending transactions")
	}

	lastBlock := bc.Blocks[len(bc.Blocks)-1]

	// Mine only up to MaxBlockSize transactions.
	transactions := bc.PendingTransactions

	if len(transactions) > bc.MaxBlockSize {
		transactions = transactions[:bc.MaxBlockSize]
	}

	block := NewBlock(
		len(bc.Blocks),
		transactions,
		lastBlock.Hash,
	)

	start := time.Now()

	result, err := miner.Mine(
		bc.Difficulty,
		func(nonce uint64) (string, error) {
			return block.CalculateHashWithNonce(nonce)
		},
	)

	if err != nil {
		return err
	}

	block.Nonce = result.Nonce
	block.Hash = result.Hash

	elapsed := time.Since(start)
	block.MiningDurationMs = elapsed.Milliseconds()

	fmt.Printf(
		"Mining completed in %.3f ms\n",
		float64(elapsed.Microseconds())/1000,
	)

	// Block is now officially mined.
	bc.Blocks = append(bc.Blocks, *block)

	// Difficulty adjustment.
	bc.adjustDifficulty()

	fmt.Printf(
		"Current Difficulty : %d\n",
		bc.Difficulty,
	)

	// Remove only the transactions that were mined.
	if len(bc.PendingTransactions) > bc.MaxBlockSize {
		bc.PendingTransactions = bc.PendingTransactions[bc.MaxBlockSize:]
	} else {
		bc.PendingTransactions = nil
	}

	return nil
}

// Print prints the blockchain.
func (bc *Blockchain) Print() {

	for _, block := range bc.Blocks {
		fmt.Println(block)
	}
}
