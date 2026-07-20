package main

import "os"

func main() {

	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {

	case "init":
		runInit()

	// case "fund":
	// 	runFund()

	// case "send":
	// 	runSend()

	// case "mine":
	// 	runMine()

	// case "validate":
	// 	runValidate()

	// case "balances":
	// 	runBalances()

	// case "print":
	// 	runPrint()

	default:
		printUsage()
	}
}