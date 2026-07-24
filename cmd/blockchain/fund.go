package main

import (
	"fmt"
	"strconv"
)

func runFund() {

	args := args()

	if len(args) != 3 {
		fmt.Println("Usage: fund <account> <amount>")
		return
	}

	account := args[1]

	amount, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		fmt.Println("Invalid amount")
		return
	}

	err = addTransaction(
		"SYSTEM",
		account,
		amount,
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Funding transaction added to pending pool.")
}
