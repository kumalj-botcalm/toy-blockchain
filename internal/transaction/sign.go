package transaction

import (
	"fmt"

	cryptoutil "github.com/kumalj-botcalm/toy-blockchain/internal/crypto"
)

// Sign signs the transaction using the supplied key pair.
func (t *Transaction) Sign(privateKey, publicKey string) error {

	signature, err := cryptoutil.Sign(
		privateKey,
		[]byte(t.ID),
	)

	if err != nil {
		return err
	}

	t.PublicKey = publicKey
	t.Signature = signature

	return nil
}

// Verify verifies the transaction signature.
func (t *Transaction) Verify() error {

	// SYSTEM transactions are trusted.
	if t.Sender == "SYSTEM" {
		return nil
	}

	if !t.IsSigned() {
		return fmt.Errorf("transaction is not signed")
	}

	valid, err := cryptoutil.Verify(
		t.PublicKey,
		[]byte(t.ID),
		t.Signature,
	)

	if err != nil {
		return err
	}

	if !valid {
		return fmt.Errorf("invalid transaction signature")
	}

	return nil
}
