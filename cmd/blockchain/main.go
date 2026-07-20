package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {

	case "init":
		fmt.Println("Initializing blockchain...")

	case "fund":
		fmt.Println("Funding account...")

	case "send":
		fmt.Println("Sending transaction...")

	case "mine":
		fmt.Println("Mining block...")

	case "validate":
		fmt.Println("Validating blockchain...")

	case "balances":
		fmt.Println("Showing balances...")

	case "print":
		fmt.Println("Printing blockchain...")

	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Toy Blockchain CLI")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  init")
	fmt.Println("  fund <account> <amount>")
	fmt.Println("  send <from> <to> <amount>")
	fmt.Println("  mine")
	fmt.Println("  validate")
	fmt.Println("  balances")
	fmt.Println("  print")
}