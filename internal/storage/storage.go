package storage

import (
	"encoding/json"
	"os"

	"github.com/kumalj-botcalm/toy-blockchain/internal/blockchain"
)

// Save writes the blockchain to a JSON file.
func Save(filename string, chain *blockchain.Blockchain) error {

	data, err := json.MarshalIndent(chain, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

// Load reads a blockchain from a JSON file.
func Load(filename string) (*blockchain.Blockchain, error) {

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var chain blockchain.Blockchain

	err = json.Unmarshal(data, &chain)
	if err != nil {
		return nil, err
	}

	return &chain, nil
}