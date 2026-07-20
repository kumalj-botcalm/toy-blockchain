package blockchain

import "github.com/kumalj-botcalm/toy-blockchain/internal/transaction"

// Block represents a single block in the blockchain.
type Block struct {
	Index        int                       `json:"index"`
	Timestamp    int64                     `json:"timestamp"`
	Transactions []transaction.Transaction `json:"transactions"`
	PreviousHash string                    `json:"previous_hash"`
	Nonce        uint64                    `json:"nonce"`
	Hash         string                    `json:"hash"`
}
