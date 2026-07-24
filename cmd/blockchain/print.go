package main

import (
	"fmt"
)

func runPrint() {

	chain, err := loadChain()
	if err != nil {
		fmt.Println("Error loading blockchain:", err)
		return
	}

	fmt.Println("========================================")

	for _, block := range chain.Blocks {

		fmt.Printf("Block #%d\n", block.Index)
		fmt.Printf("Timestamp     : %d\n", block.Timestamp)
		fmt.Printf("Hash          : %s\n", block.Hash)
		fmt.Printf("Previous Hash : %s\n", block.PreviousHash)
		fmt.Printf("Nonce         : %d\n", block.Nonce)
		fmt.Printf("Merkle Root  : %s\n", block.MerkleRoot)
		fmt.Println("Transactions:")

		if len(block.Transactions) == 0 {
			fmt.Println("  (none)")
		} else {
			for _, tx := range block.Transactions {
				fmt.Printf(
					"  %s -> %s : %.2f\n",
					tx.Sender,
					tx.Receiver,
					tx.Amount,
				)
			}
		}

		fmt.Println("========================================")
	}
}
