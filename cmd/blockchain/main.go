package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("usage:")
		fmt.Println(" init")
		fmt.Println(" add")
		fmt.Println(" mine")
		fmt.Println(" validate")
		fmt.Println(" print")
		return
	}

	switch os.Args[1] {

	case "init":
		fmt.Println("init")

	case "add":
		fmt.Println("add")

	case "mine":
		fmt.Println("mine")

	case "validate":
		fmt.Println("validate")

	case "print":
		fmt.Println("print")

	default:
		fmt.Println("unknown command")
	}
}