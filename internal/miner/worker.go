package miner

import (
	"strings"
)

// Worker searches for a valid nonce.
//
// hashFunc receives a nonce and returns the corresponding block hash.
func Worker(
	startNonce uint64,
	step uint64,
	difficulty int,
	hashFunc func(uint64) (string, error),
	result chan<- Result,
	stop <-chan struct{},
) {

	target := strings.Repeat("0", difficulty)

	nonce := startNonce

	for {

		select {

		case <-stop:
			return

		default:

			hash, err := hashFunc(nonce)
			if err == nil {

				if strings.HasPrefix(hash, target) {

					result <- Result{
						Nonce: nonce,
						Hash:  hash,
						Found: true,
					}

					return
				}
			}

			nonce += step
		}
	}
}
