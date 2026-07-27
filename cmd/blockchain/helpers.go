package main

import (
	"flag"
	"path/filepath"

	"github.com/kumalj-botcalm/toy-blockchain/internal/blockchain"
	"github.com/kumalj-botcalm/toy-blockchain/internal/storage"
	"github.com/kumalj-botcalm/toy-blockchain/internal/transaction"
)

// Default chain file.
// Can be overridden using:
// --file=mychain.json
var chainFile = flag.String(
	"file",
	blockchain.DefaultChainFile,
	"Blockchain data file",
)

func chainFilePath() string {

	path := *chainFile

	// If only a filename is provided,
	// store it inside the data folder.
	// if filepath.Dir(path) == "." {
	// 	path = filepath.Join("data", path)
	// }
	if filepath.Base(path) == path {
		path = filepath.Join("data", path)
	}

	return path
}

func loadChain() (*blockchain.Blockchain, error) {
	return storage.Load(chainFilePath())
}

func saveChain(chain *blockchain.Blockchain) error {
	return storage.Save(chainFilePath(), chain)
}

func addTransaction(sender, receiver string, amount float64) error {

	chain, err := loadChain()
	if err != nil {
		return err
	}

	tx, err := transaction.New(sender, receiver, amount)
	if err != nil {
		return err
	}

	err = chain.AddTransaction(*tx)
	if err != nil {
		return err
	}

	return saveChain(chain)
}

var interactiveArgs []string

func args() []string {
	if interactiveArgs != nil {
		return interactiveArgs
	}
	return flag.Args()
}
