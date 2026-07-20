package blockchain

import "strings"

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

		if current.PreviousHash != previous.Hash {
			return ErrInvalidPreviousHash
		}
		calculatedHash, err := current.CalculateHash()
		if err != nil {
			return err
		}

		if calculatedHash != current.Hash {
			return ErrInvalidHash
		}

		target := strings.Repeat("0", bc.Difficulty)

		if !strings.HasPrefix(current.Hash, target) {
			return ErrInvalidProofOfWork
		}

		if current.Index != previous.Index+1 {
			return ErrInvalidIndex
		}
	}
	return nil
}
