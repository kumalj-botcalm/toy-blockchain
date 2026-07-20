package blockchain

import "errors"

var (
	ErrInvalidGenesis      = errors.New("invalid genesis block")
	ErrInvalidPreviousHash = errors.New("invalid previous hash")
	ErrInvalidHash         = errors.New("invalid block hash")
	ErrInvalidProofOfWork  = errors.New("invalid proof of work")
)