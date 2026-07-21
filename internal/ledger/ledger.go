package ledger

import "fmt"

// Ledger stores account balances.
type Ledger struct {
	Balances map[string]float64
}

// New creates an empty ledger.
func New() *Ledger {
	return &Ledger{
		Balances: make(map[string]float64),
	}
}

// Balance returns the balance of an account.
func (l *Ledger) Balance(account string) float64 {

	return l.Balances[account]
}

func (l *Ledger) Credit(account string, amount float64) {

	l.Balances[account] += amount
}

func (l *Ledger) Debit(account string, amount float64) error {

	if l.Balances[account] < amount {
		return fmt.Errorf("insufficient balance")
	}

	l.Balances[account] -= amount

	return nil
}
