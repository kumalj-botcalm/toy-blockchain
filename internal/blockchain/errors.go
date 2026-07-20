package blockchain

import "errors"

var (
	ErrInvalidHash = errors.New("invalid hash")
	ErrInvalidBlock = errors.New("invalid block")
)