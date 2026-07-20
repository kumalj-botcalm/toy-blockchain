package transaction

// Pool holds pending transactions waiting to be mined.
type Pool struct {
	transactions []Transaction
}

// NewPool creates an empty transaction pool.
func NewPool() *Pool {
	return &Pool{
		transactions: make([]Transaction, 0),
	}
}

// Add adds a transaction to the pool.
func (p *Pool) Add(tx Transaction) {
	p.transactions = append(p.transactions, tx)
}

// All returns all pending transactions.
func (p *Pool) All() []Transaction {
	return p.transactions
}

// Clear removes all transactions after mining.
func (p *Pool) Clear() {
	p.transactions = make([]Transaction, 0)
}

// Count returns the number of pending transactions.
func (p *Pool) Count() int {
	return len(p.transactions)
}
