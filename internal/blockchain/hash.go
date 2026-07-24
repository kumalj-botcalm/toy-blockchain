package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

// calculateHashData is the common implementation.
func (b *Block) calculateHashData(nonce uint64) (string, error) {

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
		Nonce:        nonce,
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(bytes)

	return hex.EncodeToString(hash[:]), nil
}

// CalculateHash calculates the hash using the block's current nonce.
func (b *Block) CalculateHash() (string, error) {
	return b.calculateHashData(b.Nonce)
}

// CalculateHashWithNonce calculates the hash using a supplied nonce.
// It does NOT modify the block.
func (b *Block) CalculateHashWithNonce(nonce uint64) (string, error) {
	return b.calculateHashData(nonce)
}
