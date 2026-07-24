package blockchain

import (
	"fmt"

	"github.com/kumalj-botcalm/toy-blockchain/internal/transaction"
)

// Clone creates a deep copy of the blockchain.
func (bc *Blockchain) Clone() *Blockchain {

	if bc == nil {
		return nil
	}

	clone := &Blockchain{
		Difficulty: bc.Difficulty,
	}

	// Copy blocks
	clone.Blocks = make([]Block, len(bc.Blocks))
	copy(clone.Blocks, bc.Blocks)

	// Copy pending transactions
	clone.PendingTransactions = make([]transaction.Transaction, len(bc.PendingTransactions))
	copy(clone.PendingTransactions, bc.PendingTransactions)

	return clone
}

// ReplaceChain replaces the current chain with the candidate
// if it is valid and longer.
func (bc *Blockchain) ReplaceChain(candidate *Blockchain) error {

	if candidate == nil {
		return fmt.Errorf("candidate chain is nil")
	}

	if err := candidate.Validate(); err != nil {
		return err
	}

	if len(candidate.Blocks) <= len(bc.Blocks) {
		return fmt.Errorf("candidate chain is not longer")
	}

	replacement := candidate.Clone()

	bc.Blocks = replacement.Blocks
	bc.PendingTransactions = replacement.PendingTransactions
	bc.Difficulty = replacement.Difficulty

	return nil
}
