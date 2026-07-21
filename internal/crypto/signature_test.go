package crypto

import (
	"testing"

	"github.com/kumalj-botcalm/toy-blockchain/internal/wallet"
)

func TestSignAndVerify(t *testing.T) {

	w, err := wallet.Generate("Alice")
	if err != nil {
		t.Fatal(err)
	}

	message := []byte("Hello Blockchain")

	signature, err := Sign(w.PrivateKey, message)
	if err != nil {
		t.Fatal(err)
	}

	valid, err := Verify(
		w.PublicKey,
		message,
		signature,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !valid {
		t.Fatal("expected signature to be valid")
	}
}

func TestDetectTamperedMessage(t *testing.T) {

	w, err := wallet.Generate("Alice")
	if err != nil {
		t.Fatal(err)
	}

	message := []byte("Hello Blockchain")

	signature, err := Sign(w.PrivateKey, message)
	if err != nil {
		t.Fatal(err)
	}

	tampered := []byte("Hello Hacker")

	valid, err := Verify(
		w.PublicKey,
		tampered,
		signature,
	)

	if err != nil {
		t.Fatal(err)
	}

	if valid {
		t.Fatal("expected tampered message to fail verification")
	}
}

func TestInvalidSignature(t *testing.T) {

	w, err := wallet.Generate("Alice")
	if err != nil {
		t.Fatal(err)
	}

	message := []byte("Hello Blockchain")

	signature, err := Sign(w.PrivateKey, message)
	if err != nil {
		t.Fatal(err)
	}

	// Corrupt the signature
	signature = signature[:len(signature)-2] + "ff"

	valid, err := Verify(
		w.PublicKey,
		message,
		signature,
	)

	if err == nil && valid {
		t.Fatal("expected invalid signature")
	}
}

func TestDifferentWalletFailsVerification(t *testing.T) {

	alice, _ := wallet.Generate("Alice")
	bob, _ := wallet.Generate("Bob")

	message := []byte("Transfer 100 Coins")

	signature, err := Sign(
		alice.PrivateKey,
		message,
	)

	if err != nil {
		t.Fatal(err)
	}

	valid, err := Verify(
		bob.PublicKey,
		message,
		signature,
	)

	if err != nil {
		t.Fatal(err)
	}

	if valid {
		t.Fatal("signature should not verify using another wallet")
	}
}
