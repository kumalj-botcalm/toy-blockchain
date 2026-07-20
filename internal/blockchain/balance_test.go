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

	// Genesis funding transaction.
	tx1, _ := transaction.New("SYSTEM", "Alice", 100)

	bc.AddTransaction(*tx1)

	_ = bc.MinePendingTransactions()

	balances := bc.Balances()

	if balances["Alice"] != 100 {
		t.Fatalf("expected Alice balance to be 100")
	}
}