package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func runInteractive() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("==============================")
		fmt.Println("       TOY BLOCKCHAIN")
		fmt.Println("==============================")
		fmt.Println("1. Initialize Blockchain")
		fmt.Println("2. Create Wallet")
		fmt.Println("3. Fund Account")
		fmt.Println("4. Send Transaction")
		fmt.Println("5. Mine Block")
		fmt.Println("6. Validate Blockchain")
		fmt.Println("7. Show Balances")
		fmt.Println("8. Print Blockchain")
		fmt.Println("9. Exit")
		fmt.Println("==============================")
		fmt.Print("Enter Choice: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		
		fmt.Println()

		switch choiceStr {
		case "1":
			interactiveArgs = []string{"init"}
			runInit()
		case "2":
			fmt.Print("Enter owner name: ")
			owner, _ := reader.ReadString('\n')
			owner = strings.TrimSpace(owner)
			interactiveArgs = []string{"create-wallet", owner}
			runCreateWallet()
		case "3":
			fmt.Print("Enter account to fund: ")
			account, _ := reader.ReadString('\n')
			account = strings.TrimSpace(account)
			fmt.Print("Enter amount: ")
			amountStr, _ := reader.ReadString('\n')
			amountStr = strings.TrimSpace(amountStr)
			interactiveArgs = []string{"fund", account, amountStr}
			runFund()
		case "4":
			fmt.Print("Enter sender wallet: ")
			from, _ := reader.ReadString('\n')
			from = strings.TrimSpace(from)
			fmt.Print("Enter receiver wallet: ")
			to, _ := reader.ReadString('\n')
			to = strings.TrimSpace(to)
			fmt.Print("Enter amount: ")
			amountStr, _ := reader.ReadString('\n')
			amountStr = strings.TrimSpace(amountStr)
			interactiveArgs = []string{"send", from, to, amountStr}
			runSend()
		case "5":
			interactiveArgs = []string{"mine"}
			runMine()
		case "6":
			interactiveArgs = []string{"validate"}
			runValidate()
		case "7":
			interactiveArgs = []string{"balances"}
			runBalances()
		case "8":
			interactiveArgs = []string{"print"}
			runPrint()
		case "9":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		interactiveArgs = nil
		fmt.Println()
	}
}
