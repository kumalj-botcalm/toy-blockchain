package blockchain

import (
	"testing"

	"github.com/kumalj-botcalm/toy-blockchain/internal/transaction"
)

func TestReplaceNilChain(t *testing.T) {

	bc, err := New(2)
	if err != nil {
		t.Fatal(err)
	}

	err = bc.ReplaceChain(nil)

	if err == nil {
		t.Fatal("expected nil candidate chain to be rejected")
	}
}

func TestRejectShorterChain(t *testing.T) {

	current, _ := New(2)

	tx, _ := transaction.New(SystemAccount, "Alice", 100)
	_ = current.AddTransaction(*tx)
	_ = current.MinePendingTransactions()

	shorter, _ := New(2)

	err := current.ReplaceChain(shorter)

	if err == nil {
		t.Fatal("expected shorter chain to be rejected")
	}
}

func TestAcceptLongerChain(t *testing.T) {

	current, _ := New(2)

	longer, _ := New(2)

	tx1, _ := transaction.New(SystemAccount, "Alice", 100)
	_ = longer.AddTransaction(*tx1)
	_ = longer.MinePendingTransactions()

	tx2, _ := transaction.New(SystemAccount, "Bob", 50)
	_ = longer.AddTransaction(*tx2)
	_ = longer.MinePendingTransactions()

	err := current.ReplaceChain(longer)

	if err != nil {
		t.Fatalf("expected longer chain to replace current chain: %v", err)
	}

	if len(current.Blocks) != len(longer.Blocks) {
		t.Fatal("chain replacement failed")
	}
}

func TestRejectInvalidChain(t *testing.T) {

	current, _ := New(2)

	invalid, _ := New(2)

	tx, _ := transaction.New(SystemAccount, "Alice", 100)
	_ = invalid.AddTransaction(*tx)
	_ = invalid.MinePendingTransactions()

	// Tamper the block
	invalid.Blocks[1].PreviousHash = "tampered"

	err := current.ReplaceChain(invalid)

	if err == nil {
		t.Fatal("expected invalid chain to be rejected")
	}
}
