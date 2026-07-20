package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kumalj-botcalm/toy-blockchain/internal/blockchain"
)

func runFund() {

	if len(os.Args) != 4 {
		fmt.Println("Usage:")
		fmt.Println("  fund <account> <amount>")
		return
	}

	account := os.Args[2]

	amount, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Invalid amount")
		return
	}

	err = addTransaction(
		blockchain.SystemAccount,
		account,
		amount,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Funding transaction added to pending pool.")
}