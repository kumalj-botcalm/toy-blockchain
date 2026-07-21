package merkle

import (
	"testing"

	"github.com/kumalj-botcalm/toy-blockchain/internal/transaction"
)

func TestEmptyMerkleRoot(t *testing.T) {

	root, err := BuildMerkleRoot([]transaction.Transaction{})
	if err != nil {
		t.Fatal(err)
	}

	expected := HashLeaf([]byte{})

	if root != expected {
		t.Fatalf("expected %s, got %s", expected, root)
	}
}

func TestSingleTransactionMerkleRoot(t *testing.T) {

	tx, err := transaction.New("Alice", "Bob", 100)
	if err != nil {
		t.Fatal(err)
	}

	root, err := BuildMerkleRoot([]transaction.Transaction{*tx})
	if err != nil {
		t.Fatal(err)
	}

	if root == "" {
		t.Fatal("expected non-empty merkle root")
	}
}

func TestTwoTransactionMerkleRoot(t *testing.T) {

	tx1, err := transaction.New("Alice", "Bob", 100)
	if err != nil {
		t.Fatal(err)
	}

	tx2, err := transaction.New("Bob", "Charlie", 50)
	if err != nil {
		t.Fatal(err)
	}

	root, err := BuildMerkleRoot([]transaction.Transaction{
		*tx1,
		*tx2,
	})
	if err != nil {
		t.Fatal(err)
	}

	if root == "" {
		t.Fatal("expected non-empty merkle root")
	}
}

func TestMerkleRootDeterministic(t *testing.T) {

	tx1, err := transaction.New("Alice", "Bob", 100)
	if err != nil {
		t.Fatal(err)
	}

	tx2, err := transaction.New("Bob", "Charlie", 50)
	if err != nil {
		t.Fatal(err)
	}

	transactions := []transaction.Transaction{
		*tx1,
		*tx2,
	}

	root1, err := BuildMerkleRoot(transactions)
	if err != nil {
		t.Fatal(err)
	}

	root2, err := BuildMerkleRoot(transactions)
	if err != nil {
		t.Fatal(err)
	}

	if root1 != root2 {
		t.Fatal("merkle root should be deterministic")
	}
}

func TestMerkleRootChangesWhenTransactionChanges(t *testing.T) {

	tx1, err := transaction.New("Alice", "Bob", 100)
	if err != nil {
		t.Fatal(err)
	}

	tx2, err := transaction.New("Bob", "Charlie", 50)
	if err != nil {
		t.Fatal(err)
	}

	transactions := []transaction.Transaction{
		*tx1,
		*tx2,
	}

	root1, err := BuildMerkleRoot(transactions)
	if err != nil {
		t.Fatal(err)
	}

	transactions[1].Amount = 999

	root2, err := BuildMerkleRoot(transactions)
	if err != nil {
		t.Fatal(err)
	}

	if root1 == root2 {
		t.Fatal("merkle root should change when transaction data changes")
	}
}