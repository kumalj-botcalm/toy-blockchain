package main

import (
	"fmt"

	"github.com/kumalj-botcalm/toy-blockchain/internal/blockchain"
)

func runInit() {

	fmt.Println("Initializing blockchain...")

	bc, err := blockchain.New(blockchain.DefaultDifficulty)
	if err != nil {
		fmt.Println("Error initializing blockchain:", err)
		return
	}

	err = saveChain(bc)
	if err != nil {
		fmt.Println("Error saving blockchain:", err)
		return
	}

	fmt.Println("Blockchain initialized.")
}
