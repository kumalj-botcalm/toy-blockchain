package blockchain

import "errors"

var (
	ErrInvalidGenesis      = errors.New("invalid genesis block")
	ErrInvalidPreviousHash = errors.New("invalid previous hash")
	ErrInvalidHash         = errors.New("invalid block hash")
	ErrInvalidProofOfWork  = errors.New("invalid proof of work")
	ErrInvalidIndex        = errors.New("invalid block index")
	ErrInvalidTimestamp    = errors.New("invalid timestamp")
	ErrInvalidMerkleRoot   = errors.New("invalid merkle root")
)
