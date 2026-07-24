package blockchain

import (
	"strings"

	"github.com/kumalj-botcalm/toy-blockchain/internal/merkle"
)

func (bc *Blockchain) Validate() error {

	if len(bc.Blocks) == 0 {
		return ErrInvalidGenesis
	}

	genesis := bc.Blocks[0]

	expectedGenesis, err := NewGenesisBlock()
	if err != nil {
		return err
	}

	if genesis.Hash != expectedGenesis.Hash {
		return ErrInvalidGenesis
	}

	for i := 1; i < len(bc.Blocks); i++ {

		current := bc.Blocks[i]
		previous := bc.Blocks[i-1]

		// Previous hash validation
		if current.PreviousHash != previous.Hash {
			return ErrInvalidPreviousHash
		}

		// Merkle Root validation
		expectedRoot, err := merkle.BuildMerkleRoot(current.Transactions)
		if err != nil {
			return err
		}

		if expectedRoot != current.MerkleRoot {
			return ErrInvalidMerkleRoot
		}

		// Transaction signature validation
		for _, tx := range current.Transactions {

			if tx.Sender == SystemAccount {
				continue
			}

			if err := tx.Verify(); err != nil {
				return err
			}

		}

		// Block hash validation
		calculatedHash, err := current.CalculateHash()
		if err != nil {
			return err
		}

		if calculatedHash != current.Hash {
			return ErrInvalidHash
		}

		// Proof-of-Work validation
		expectedDifficulty := ExpectedDifficultyAt(bc.Blocks, i)
		target := strings.Repeat("0", expectedDifficulty)

		if !strings.HasPrefix(current.Hash, target) {
			return ErrInvalidProofOfWork
		}

		// Index validation
		if current.Index != previous.Index+1 {
			return ErrInvalidIndex
		}

		// Timestamp validation
		if current.Timestamp < previous.Timestamp {
			return ErrInvalidTimestamp
		}
	}

	return nil
}
