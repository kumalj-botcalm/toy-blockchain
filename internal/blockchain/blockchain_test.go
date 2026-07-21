package blockchain

import "testing"

func TestBlockchainStartsWithGenesisBlock(t *testing.T) {

	bc, err := New(3)
	if err != nil {
		t.Fatal(err)
	}

	if len(bc.Blocks) != 1 {
		t.Fatalf("expected 1 genesis block")
	}

	if bc.Blocks[0].Index != 0 {
		t.Fatalf("expected genesis index to be zero")
	}

	if bc.Blocks[0].PreviousHash != GenesisPreviousHash {
		t.Fatalf("unexpected genesis previous hash")
	}
}
