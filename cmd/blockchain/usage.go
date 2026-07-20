package main

import "fmt"

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