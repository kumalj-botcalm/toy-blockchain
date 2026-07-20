package main

import (
	"github.com/kumalj-botcalm/toy-blockchain/internal/blockchain"
	"github.com/kumalj-botcalm/toy-blockchain/internal/storage"
	"github.com/kumalj-botcalm/toy-blockchain/internal/transaction"
)

func loadChain() (*blockchain.Blockchain, error) {
	return storage.Load(blockchain.DefaultChainFile)
}

func saveChain(chain *blockchain.Blockchain) error {
	return storage.Save(blockchain.DefaultChainFile, chain)
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