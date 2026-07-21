package merkle

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashLeaf hashes a single leaf value.
func HashLeaf(data []byte) string {

	hash := sha256.Sum256(data)

	return hex.EncodeToString(hash[:])
}

// HashPair hashes two child hashes together.
func HashPair(left, right string) string {

	data := []byte(left + right)

	hash := sha256.Sum256(data)

	return hex.EncodeToString(hash[:])
}
