package blockchain

import "testing"

func TestBlockString(t *testing.T) {

	bc, err := New(2)
	if err != nil {
		t.Fatal(err)
	}

	if bc.Blocks[0].String() == "" {
		t.Fatal("expected block string")
	}
}