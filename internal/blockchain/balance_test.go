package blockchain

import (
	"testing"

	"github.com/kumalj-botcalm/toy-blockchain/internal/transaction"
)

func TestBalances(t *testing.T) {

	bc, err := New(2)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := transaction.New("SYSTEM", "Alice", 100)
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

	balances := bc.Balances()

	if balances["Alice"] != 100 {
		t.Fatalf("expected Alice balance to be 100, got %v", balances["Alice"])
	}
}

func TestRejectOverspending(t *testing.T) {

	bc, _ := New(2)

	tx, _ := transaction.New("Alice", "Bob", 100)

	err := bc.AddTransaction(*tx)

	if err == nil {
		t.Fatal("expected overspending transaction to be rejected")
	}
}

func TestSystemFunding(t *testing.T) {

	bc, _ := New(2)

	tx, _ := transaction.New(SystemAccount, "Alice", 100)

	err := bc.AddTransaction(*tx)

	if err != nil {
		t.Fatal(err)
	}
}
