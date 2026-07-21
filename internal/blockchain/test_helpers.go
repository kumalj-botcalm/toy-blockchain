package blockchain

import (
	"testing"

	"github.com/kumalj-botcalm/toy-blockchain/internal/transaction"
)

func fundAccount(t *testing.T, bc *Blockchain, account string, amount float64) {

	t.Helper()

	tx, err := transaction.New(SystemAccount, account, amount)
	if err != nil {
		t.Fatal(err)
	}

	err = bc.AddTransaction(*tx)
	if err != nil {
		t.Fatal(err)
	}

	err = bc.MinePendingTransactions()
	if err != nil {
		t.Fatal(err)
	}
}