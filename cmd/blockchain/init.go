package main

import (
	"fmt"
	"github.com/kumalj-botcalm/toy-blockchain/internal/blockchain"
	"github.com/kumalj-botcalm/toy-blockchain/internal/storage"
)

func runInit() {
	fmt.Println("Initializing blockchain...")

	bc, err := blockchain.New(blockchain.DefaultDifficulty)
	if err != nil {
		fmt.Println("Error initializing blockchain:", err)
		return
	}

	err = storage.Save(blockchain.DefaultChainFile, bc)
	if err != nil {
		fmt.Println("Error saving blockchain:", err)
		return
	}

	fmt.Println("Blockchain initialized.")
}
