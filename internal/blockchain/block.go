package blockchain

import (
	"fmt"
	"strings"
	"time"

	"github.com/kumalj-botcalm/toy-blockchain/internal/merkle"
	"github.com/kumalj-botcalm/toy-blockchain/internal/transaction"
)

// Block represents a single block in the blockchain.
type Block struct {
	Index        int                       `json:"index"`
	Timestamp    int64                     `json:"timestamp"`
	Transactions []transaction.Transaction `json:"transactions"`

	MerkleRoot string `json:"merkle_root"`

	PreviousHash string `json:"previous_hash"`

	Nonce uint64 `json:"nonce"`

	// Analytics only.
	MiningDurationMs int64 `json:"mining_duration_ms"`

	Hash string `json:"hash"`
}

// NewBlock creates a new block.
func NewBlock(
	index int,
	transactions []transaction.Transaction,
	previousHash string,
) *Block {

	block := &Block{
		Index:        index,
		Timestamp:    time.Now().Unix(),
		Transactions: transactions,
		PreviousHash: previousHash,
		Nonce:        0,
	}

	root, err := merkle.BuildMerkleRoot(transactions)
	if err != nil {
		return nil
	}

	block.MerkleRoot = root
	return block
}

// String returns a formatted representation of the block.
func (b Block) String() string {

	var out strings.Builder

	out.WriteString("----------------------------------------\n")

	out.WriteString(fmt.Sprintf("Block #%d\n", b.Index))

	out.WriteString(fmt.Sprintf("Hash: %s\n", b.Hash))

	out.WriteString(fmt.Sprintf("Previous: %s\n", b.PreviousHash))

	out.WriteString(fmt.Sprintf("Nonce: %d\n", b.Nonce))

	out.WriteString(fmt.Sprintf("Timestamp: %s\n",
		time.Unix(b.Timestamp, 0).Format(time.RFC3339)))

	out.WriteString("Transactions:\n")

	out.WriteString(fmt.Sprintf("Merkle Root: %s\n", b.MerkleRoot))

	if len(b.Transactions) == 0 {
		out.WriteString("  (none)\n")
	} else {
		for _, tx := range b.Transactions {
			out.WriteString("  ")
			out.WriteString(tx.String())
			out.WriteString("\n")
		}
	}

	return out.String()
}
