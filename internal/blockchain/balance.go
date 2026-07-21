package blockchain

// Balances rebuilds account balances from all mined blocks.
func (bc *Blockchain) Balances() map[string]float64 {

	balances := make(map[string]float64)

	for _, block := range bc.Blocks {

		for _, tx := range block.Transactions {

			if tx.Sender != SystemAccount {
				balances[tx.Sender] -= tx.Amount
			}

			balances[tx.Receiver] += tx.Amount
		}
	}

	return balances
}
