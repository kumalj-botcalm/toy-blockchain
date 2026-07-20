package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

// CalculateHash calculates the SHA-256 hash of a block.
//
// IMPORTANT:
// The Hash field itself is NOT included in the hash calculation.
// Otherwise every recalculation would generate a different value.
func (b *Block) CalculateHash() (string, error) {

	type blockForHash struct {
		Index        int
		Timestamp    int64
		Transactions interface{}
		PreviousHash string
		Nonce        uint64
	}

	data := blockForHash{
		Index:        b.Index,
		Timestamp:    b.Timestamp,
		Transactions: b.Transactions,
		PreviousHash: b.PreviousHash,
		Nonce:        b.Nonce,
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(bytes)

	return hex.EncodeToString(hash[:]), nil
}