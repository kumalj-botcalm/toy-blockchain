package blockchain

import (
	"fmt"
	"time"
)

const (
	AdjustmentInterval = 5

	MinimumDifficulty = 1
)

// adjustDifficulty adjusts mining difficulty based on
// average block generation time.
func (bc *Blockchain) adjustDifficulty() {

	// Need at least AdjustmentInterval + 1 blocks
	// to calculate AdjustmentInterval block times.
	if len(bc.Blocks) < AdjustmentInterval+1 {
		return
	}

	start := len(bc.Blocks) - AdjustmentInterval - 1

	var total time.Duration

	for i := start + 1; i < len(bc.Blocks); i++ {

		prev := bc.Blocks[i-1]
		curr := bc.Blocks[i]

		total += time.Duration(
			curr.Timestamp-prev.Timestamp,
		) * time.Second
	}

	average := total / AdjustmentInterval

	oldDifficulty := bc.Difficulty

	switch {

	case average < time.Second:

		bc.Difficulty++

	case average > 3*time.Second:

		if bc.Difficulty > MinimumDifficulty {
			bc.Difficulty--
		}
	}

	if oldDifficulty != bc.Difficulty {

		fmt.Printf("\nDifficulty adjusted!\n")
		fmt.Printf(
			"%d -> %d\n\n",
			oldDifficulty,
			bc.Difficulty,
		)
	}
}
