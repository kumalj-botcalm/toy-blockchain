package blockchain

import (
	"os"
	"testing"

	"github.com/kumalj-botcalm/toy-blockchain/internal/transaction"
	"github.com/kumalj-botcalm/toy-blockchain/internal/wallet"
)

func createSignedTransaction(
	t *testing.T,
	sender, receiver string,
	amount float64,
) *transaction.Transaction {

	t.Helper()

	tx, err := transaction.New(sender, receiver, amount)
	if err != nil {
		t.Fatal(err)
	}

	w, err := wallet.Generate(sender)
	if err != nil {
		t.Fatal(err)
	}

	err = w.Save()
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		_ = os.Remove("wallets/" + sender + ".json")
	})

	err = tx.Sign(
		w.PrivateKey,
		w.PublicKey,
	)

	if err != nil {
		t.Fatal(err)
	}

	return tx
}

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

	// Signed transaction
	tx := createSignedTransaction(
		t,
		"Alice",
		"Bob",
		50,
	)

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

	// Signed transaction
	tx := createSignedTransaction(
		t,
		"Alice",
		"Bob",
		50,
	)

	err = bc.AddTransaction(*tx)
	if err != nil {
		t.Fatal(err)
	}

	err = bc.MinePendingTransactions()
	if err != nil {
		t.Fatal(err)
	}

	// Tamper transaction
	bc.Blocks[2].Transactions[0].Amount = 999999

	err = bc.Validate()

	if err == nil {
		t.Fatal("expected tampered chain to be invalid")
	}
}

func TestInvalidBlockIndex(t *testing.T) {

	bc, err := New(2)
	if err != nil {
		t.Fatal(err)
	}

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

	bc.Blocks[1].Index = 99

	err = bc.Validate()

	if err == nil {
		t.Fatal("expected invalid block index")
	}
}

func TestInvalidTimestamp(t *testing.T) {

	bc, err := New(2)
	if err != nil {
		t.Fatal(err)
	}

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

	bc.Blocks[1].Timestamp = bc.Blocks[0].Timestamp - 1

	err = bc.Validate()

	if err == nil {
		t.Fatal("expected invalid timestamp")
	}
}
