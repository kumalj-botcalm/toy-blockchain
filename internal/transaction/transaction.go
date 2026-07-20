package transaction

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

// Transaction represents a transfer of funds between two accounts.
type Transaction struct {
	ID        string  `json:"id"`
	Sender    string  `json:"sender"`
	Receiver  string  `json:"receiver"`
	Amount    float64 `json:"amount"`
	Timestamp int64   `json:"timestamp"`
}

// New creates a new transaction.
func New(sender, receiver string, amount float64) (*Transaction, error) {
	if sender == "" {
		return nil, fmt.Errorf("sender cannot be empty")
	}

	if receiver == "" {
		return nil, fmt.Errorf("receiver cannot be empty")
	}

	if amount <= 0 {
		return nil, fmt.Errorf("amount must be greater than zero")
	}

	tx := &Transaction{
		Sender:    sender,
		Receiver:  receiver,
		Amount:    amount,
		Timestamp: time.Now().Unix(),
	}

	id, err := tx.GenerateID()
	if err != nil {
		return nil, err
	}

	tx.ID = id

	return tx, nil
}

// GenerateID creates a deterministic SHA-256 ID for the transaction.
func (t *Transaction) GenerateID() (string, error) {

	data, err := json.Marshal(struct {
		Sender    string
		Receiver  string
		Amount    float64
		Timestamp int64
	}{
		Sender:    t.Sender,
		Receiver:  t.Receiver,
		Amount:    t.Amount,
		Timestamp: t.Timestamp,
	})

	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(data)

	return hex.EncodeToString(hash[:]), nil
}

// String returns a human-readable representation of the transaction.
func (t Transaction) String() string {
	return fmt.Sprintf(
		"%s -> %s : %.2f",
		t.Sender,
		t.Receiver,
		t.Amount,
	)
}
