package miner

import (
	"strings"
	"testing"
)

func TestMine(t *testing.T) {

	result, err := Mine(
		2,
		func(nonce uint64) (string, error) {

			if nonce == 100 {

				return "00abcdef", nil
			}

			return "abcdef", nil
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !result.Found {
		t.Fatal("expected miner to find a result")
	}

	if result.Nonce != 100 {
		t.Fatalf("expected nonce 100, got %d", result.Nonce)
	}

	if !strings.HasPrefix(result.Hash, "00") {
		t.Fatal("expected hash with required difficulty")
	}
}
