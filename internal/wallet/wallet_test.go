package wallet

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateWallet(t *testing.T) {

	w, err := Generate("Alice")
	if err != nil {
		t.Fatal(err)
	}

	if w.Owner != "Alice" {
		t.Fatalf("expected owner Alice, got %s", w.Owner)
	}

	if w.PublicKey == "" {
		t.Fatal("public key should not be empty")
	}

	if w.PrivateKey == "" {
		t.Fatal("private key should not be empty")
	}
}

func TestGenerateWalletEmptyOwner(t *testing.T) {

	_, err := Generate("")

	if err == nil {
		t.Fatal("expected error for empty owner")
	}
}

func TestSaveAndLoadWallet(t *testing.T) {

	w, err := Generate("TestUser")
	if err != nil {
		t.Fatal(err)
	}

	err = w.Save()
	if err != nil {
		t.Fatal(err)
	}

	loaded, err := Load("TestUser")
	if err != nil {
		t.Fatal(err)
	}

	if loaded.Owner != w.Owner {
		t.Fatal("owner mismatch")
	}

	if loaded.PublicKey != w.PublicKey {
		t.Fatal("public key mismatch")
	}

	if loaded.PrivateKey != w.PrivateKey {
		t.Fatal("private key mismatch")
	}

	// cleanup
	filename := filepath.Join(WalletDirectory, "TestUser.json")
	_ = os.Remove(filename)
}