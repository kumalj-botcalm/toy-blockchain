package miner

// Result represents a successful mining result.
type Result struct {
	Nonce uint64
	Hash  string
	Found bool
}
