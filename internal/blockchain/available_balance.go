package blockchain

// AvailableBalances returns balances after applying pending transactions.
func (bc *Blockchain) AvailableBalances() map[string]float64 {
	balances := bc.Balances()

	copyBalances := make(map[string]float64)

	for account, balance := range balances {
		copyBalances[account] = balance
	}

	for _, tx := range bc.PendingTransactions {

		if tx.Sender != SystemAccount {

			copyBalances[tx.Sender] -= tx.Amount
		}

		copyBalances[tx.Receiver] += tx.Amount
	}

	return copyBalances

}
