package main

import (
	"fmt"

	"github.com/kumalj-botcalm/toy-blockchain/internal/wallet"
)

func runCreateWallet() {

	args := args()

	if len(args) != 2 {
		fmt.Println("Usage:")
		fmt.Println("  create-wallet <owner>")
		return
	}

	owner := args[1]

	fmt.Println("Creating wallet...")

	w, err := wallet.Generate(owner)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = w.Save()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println()
	fmt.Println("Wallet created successfully.")
	fmt.Println()
	fmt.Println("Owner      :", w.Owner)
	key := w.PublicKey
	if len(key) > 32 {
		key = key[:32] + "..."
	}

	fmt.Println("Public Key :", key)
	fmt.Printf("Saved To   : wallets/%s.json\n", owner)
}
