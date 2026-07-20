package main

import (
	"fmt"
	"os"
	"strconv"
)

func runSend() {

	if len(os.Args) != 5 {
		fmt.Println("Usage:")
		fmt.Println("  send <sender> <receiver> <amount>")
		return
	}

	sender := os.Args[2]
	receiver := os.Args[3]

	amount, err := strconv.ParseFloat(os.Args[4], 64)
	if err != nil {
		fmt.Println("Invalid amount")
		return
	}

	err = addTransaction(
		sender,
		receiver,
		amount,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Transaction added to pending pool.")
}