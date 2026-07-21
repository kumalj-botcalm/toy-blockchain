package main

import "fmt"

func runMine() {

	chain, err := loadChain()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Mining pending transactions...")

	err = chain.MinePendingTransactions()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = saveChain(chain)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	last := chain.Blocks[len(chain.Blocks)-1]

	fmt.Println()
	fmt.Println("Block mined successfully.")
	fmt.Printf("Height : %d\n", last.Index)
	fmt.Printf("Hash   : %s\n", last.Hash)
	fmt.Printf("Nonce  : %d\n", last.Nonce)
}
