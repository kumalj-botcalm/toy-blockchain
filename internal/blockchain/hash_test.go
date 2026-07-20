package blockchain

import "testing"

func TestHashIsDeterministic(t *testing.T) {

	block := Block{
		Index:        1,
		Timestamp:    1721174400,
		PreviousHash: "0000",
		Nonce:        10,
	}

	hash1, err := block.CalculateHash()
	if err != nil {
		t.Fatal(err)
	}

	hash2, err := block.CalculateHash()
	if err != nil {
		t.Fatal(err)
	}

	if hash1 != hash2 {
		t.Fatalf("expected identical hashes")
	}
}