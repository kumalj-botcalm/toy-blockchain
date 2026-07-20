package storage

import (
	"os"
	"testing"

	"github.com/kumalj-botcalm/toy-blockchain/internal/blockchain"
)

func TestSaveAndLoad(t *testing.T) {

	filename := "test_chain.json"

	bc, err := blockchain.New(2)
	if err != nil {
		t.Fatal(err)
	}

	err = Save(filename, bc)
	if err != nil {
		t.Fatal(err)
	}

	loaded, err := Load(filename)
	if err != nil {
		t.Fatal(err)
	}

	if len(loaded.Blocks) != len(bc.Blocks) {
		t.Fatal("block count mismatch")
	}

	_ = os.Remove(filename)
}