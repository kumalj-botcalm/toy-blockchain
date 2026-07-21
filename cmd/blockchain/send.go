package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kumalj-botcalm/toy-blockchain/internal/transaction"
	"github.com/kumalj-botcalm/toy-blockchain/internal/wallet"
)

func runSend() {

	if len(os.Args) != 5 {
		fmt.Println("Usage:")
		fmt.Println("  send <from> <to> <amount>")
		return
	}

	from := os.Args[2]
	to := os.Args[3]

	amount, err := strconv.ParseFloat(os.Args[4], 64)
	if err != nil {
		fmt.Println("Invalid amount")
		return
	}

	chain, err := loadChain()
	if err != nil {
		fmt.Println("Error loading blockchain:", err)
		return
	}

	tx, err := transaction.New(from, to, amount)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Load sender wallet
	w, err := wallet.Load(from)
	if err != nil {
		fmt.Println("Error loading wallet:", err)
		return
	}

	// Sign transaction
	err = tx.Sign(
		w.PrivateKey,
		w.PublicKey,
	)
	if err != nil {
		fmt.Println("Error signing transaction:", err)
		return
	}

	// Add transaction
	err = chain.AddTransaction(*tx)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Save blockchain
	err = saveChain(chain)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Signed transaction added to pending pool.")
}
