package main

import (
	"fmt"
	"sort"
)

func runBalances() {

	chain, err := loadChain()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	balances := chain.Balances()

	if len(balances) == 0 {
		fmt.Println("No balances.")
		return
	}

	// Sort account names for deterministic output
	accounts := make([]string, 0, len(balances))
	for account := range balances {
		accounts = append(accounts, account)
	}

	sort.Strings(accounts)

	fmt.Println("Account Balances")
	fmt.Println("----------------")

	for _, account := range accounts {
		fmt.Printf("%-10s : %.2f\n", account, balances[account])
	}
}
