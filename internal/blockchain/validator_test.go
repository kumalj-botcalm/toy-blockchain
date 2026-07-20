package blockchain

import (
	"testing"

	"github.com/kumalj-botcalm/toy-blockchain/internal/transaction"
)

func TestValidateValidBlockchain(t *testing.T) {

	bc, err := New(2)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := transaction.New("Alice", "Bob", 100)
	if err != nil {
		t.Fatal(err)
	}

	bc.AddTransaction(*tx)

	err = bc.MinePendingTransactions()
	if err != nil {
		t.Fatal(err)
	}

	err = bc.Validate()
	if err != nil {
		t.Fatalf("expected blockchain to be valid: %v", err)
	}
}

func TestDetectTamperedBlock(t *testing.T) {

	bc, _ := New(2)

	tx, _ := transaction.New("Alice", "Bob", 100)

	bc.AddTransaction(*tx)

	_ = bc.MinePendingTransactions()

	bc.Blocks[1].Transactions[0].Amount = 999999

	err := bc.Validate()

	if err == nil {
		t.Fatal("expected tampered chain to be invalid")
	}
}