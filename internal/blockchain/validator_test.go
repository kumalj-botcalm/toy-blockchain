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

	// Fund Alice
	fund, err := transaction.New(SystemAccount, "Alice", 100)
	if err != nil {
		t.Fatal(err)
	}

	err = bc.AddTransaction(*fund)
	if err != nil {
		t.Fatal(err)
	}

	err = bc.MinePendingTransactions()
	if err != nil {
		t.Fatal(err)
	}

	// Alice spends money
	tx, err := transaction.New("Alice", "Bob", 50)
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

	err = bc.Validate()
	if err != nil {
		t.Fatalf("expected blockchain to be valid: %v", err)
	}
}

func TestDetectTamperedBlock(t *testing.T) {

	bc, _ := New(2)

	// Fund Alice
	fund, _ := transaction.New(SystemAccount, "Alice", 100)

	err := bc.AddTransaction(*fund)
	if err != nil {
		t.Fatal(err)
	}

	err = bc.MinePendingTransactions()
	if err != nil {
		t.Fatal(err)
	}

	// Alice spends
	tx, _ := transaction.New("Alice", "Bob", 50)

	err = bc.AddTransaction(*tx)
	if err != nil {
		t.Fatal(err)
	}

	err = bc.MinePendingTransactions()
	if err != nil {
		t.Fatal(err)
	}

	// Tamper second mined block
	bc.Blocks[2].Transactions[0].Amount = 999999

	err = bc.Validate()

	if err == nil {
		t.Fatal("expected tampered chain to be invalid")
	}
}


func TestInvalidBlockIndex(t *testing.T) {

	bc, _ := New(2)

	fund, _ := transaction.New(SystemAccount, "Alice", 100)
	_ = bc.AddTransaction(*fund)
	_ = bc.MinePendingTransactions()

	bc.Blocks[1].Index = 99

	err := bc.Validate()

	if err == nil {
		t.Fatal("expected invalid block index")
	}
}