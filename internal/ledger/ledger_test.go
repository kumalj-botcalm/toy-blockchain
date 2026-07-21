package ledger

import "testing"

func TestLedger(t *testing.T) {

	l := New()

	l.Credit("Alice", 100)

	if l.Balance("Alice") != 100 {
		t.Fatal("wrong balance")
	}

	err := l.Debit("Alice", 50)
	if err != nil {
		t.Fatal(err)
	}

	if l.Balance("Alice") != 50 {
		t.Fatal("wrong balance")
	}
}
