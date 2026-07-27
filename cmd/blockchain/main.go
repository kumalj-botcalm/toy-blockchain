package main

import "flag"

func main() {

	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		runInteractive()
		return
	}

	switch args[0] {

	case "init":
		runInit()

	case "fund":
		runFund()

	case "send":
		runSend()

	case "mine":
		runMine()

	case "validate":
		runValidate()

	case "balances":
		runBalances()

	case "print":
		runPrint()

	case "create-wallet":
		runCreateWallet()

	default:
		printUsage()
	}
}
