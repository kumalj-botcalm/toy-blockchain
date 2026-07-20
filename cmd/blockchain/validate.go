package main

import "fmt"

func runValidate() {

	chain, err := loadChain()
	if err != nil {
		fmt.Println("Error loading blockchain:", err)
		return
	}

	err = chain.Validate()
	if err != nil {
		fmt.Println("Blockchain is invalid:")
		fmt.Println(err)
		return
	}

	fmt.Println("Blockchain is valid.")
}