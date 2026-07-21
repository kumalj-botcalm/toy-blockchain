package blockchain

import (
	"fmt"
	"strings"

	"github.com/kumalj-botcalm/toy-blockchain/internal/merkle"
)

func (bc *Blockchain) Validate() error {

	if len(bc.Blocks) == 0 {
		return ErrInvalidGenesis
	}

	genesis := bc.Blocks[0]

	if genesis.Index != 0 {
		return ErrInvalidGenesis
	}

	if genesis.PreviousHash != GenesisPreviousHash {
		return ErrInvalidGenesis
	}

	for i := 1; i < len(bc.Blocks); i++ {

		current := bc.Blocks[i]
		previous := bc.Blocks[i-1]

		// 1. Previous hash
		if current.PreviousHash != previous.Hash {
			return ErrInvalidPreviousHash
		}

		// 2. Rebuild Merkle Root
		expectedRoot, err := merkle.BuildMerkleRoot(current.Transactions)
		if err != nil {
			return err
		}

		if expectedRoot != current.MerkleRoot {
			fmt.Println("Stored Root  :", current.MerkleRoot)
			fmt.Println("Expected Root:", expectedRoot)
			return ErrInvalidMerkleRoot
		}

		// 3. Recalculate block hash
		calculatedHash, err := current.CalculateHash()
		if err != nil {
			return err
		}

		if calculatedHash != current.Hash {
			return ErrInvalidHash
		}

		// 4. Proof of Work
		target := strings.Repeat("0", bc.Difficulty)
		if !strings.HasPrefix(current.Hash, target) {
			return ErrInvalidProofOfWork
		}

		// 5. Height
		if current.Index != previous.Index+1 {
			return ErrInvalidIndex
		}

		// 6. Timestamp
		if current.Timestamp < previous.Timestamp {
			return ErrInvalidTimestamp
		}

	}
	return nil
}
