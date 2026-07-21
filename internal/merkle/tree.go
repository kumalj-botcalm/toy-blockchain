package merkle

import (
	"encoding/json"

	"github.com/kumalj-botcalm/toy-blockchain/internal/transaction"
)

// BuildMerkleRoot builds a Merkle Tree and returns the Merkle Root.
func BuildMerkleRoot(transactions []transaction.Transaction) (string, error) {

	// Empty tree
	if len(transactions) == 0 {
		return HashLeaf([]byte{}), nil
	}

	// Create leaf hashes
	var hashes []string

	for _, tx := range transactions {

		data, err := json.Marshal(tx)
		if err != nil {
			return "", err
		}

		hashes = append(hashes, HashLeaf(data))
	}

	// Build tree until only one hash remains
	for len(hashes) > 1 {

		var nextLevel []string

		for i := 0; i < len(hashes); i += 2 {

			left := hashes[i]

			right := left

			// If there is a right node, use it.
			// Otherwise duplicate the left node.
			if i+1 < len(hashes) {
				right = hashes[i+1]
			}

			parent := HashPair(left, right)

			nextLevel = append(nextLevel, parent)
		}

		hashes = nextLevel
	}

	return hashes[0], nil
}