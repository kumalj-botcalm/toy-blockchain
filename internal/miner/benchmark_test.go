package miner

import (
	"testing"
)

func BenchmarkMineDifficulty2(b *testing.B) {

	for i := 0; i < b.N; i++ {

		_, err := Mine(
			2,
			func(nonce uint64) (string, error) {

				// Simulate a valid hash after some work.
				if nonce == 1000 {
					return "00abcdef1234567890", nil
				}

				return "abcdef1234567890", nil
			},
		)

		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMineDifficulty3(b *testing.B) {

	for i := 0; i < b.N; i++ {

		_, err := Mine(
			3,
			func(nonce uint64) (string, error) {

				if nonce == 5000 {
					return "000abcdef1234567890", nil
				}

				return "abcdef1234567890", nil
			},
		)

		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMineDifficulty4(b *testing.B) {

	for i := 0; i < b.N; i++ {

		_, err := Mine(
			4,
			func(nonce uint64) (string, error) {

				if nonce == 15000 {
					return "0000abcdef1234567890", nil
				}

				return "abcdef1234567890", nil
			},
		)

		if err != nil {
			b.Fatal(err)
		}
	}
}
