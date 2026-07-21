package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

// CalculateHash calculates the SHA-256 hash of a block.
//
// The block hash is calculated WITHOUT the Hash field itself.
// Transactions are represented only by the Merkle Root.
func (b *Block) CalculateHash() (string, error) {

	type blockForHash struct {
		Index        int
		Timestamp    int64
		MerkleRoot   string
		PreviousHash string
		Nonce        uint64
	}

	data := blockForHash{
		Index:        b.Index,
		Timestamp:    b.Timestamp,
		MerkleRoot:   b.MerkleRoot,
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