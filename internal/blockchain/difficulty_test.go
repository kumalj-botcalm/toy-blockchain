package blockchain

import (
	"testing"
)

func TestDifficultyNeverBelowMinimum(t *testing.T) {

	bc, err := New(MinimumDifficulty)
	if err != nil {
		t.Fatal(err)
	}

	bc.Difficulty = MinimumDifficulty

	bc.adjustDifficulty()

	if bc.Difficulty < MinimumDifficulty {
		t.Fatalf(
			"difficulty should never be below %d",
			MinimumDifficulty,
		)
	}
}

func TestDifficultyWithoutEnoughBlocks(t *testing.T) {

	bc, err := New(2)
	if err != nil {
		t.Fatal(err)
	}

	original := bc.Difficulty

	bc.adjustDifficulty()

	if bc.Difficulty != original {
		t.Fatal("difficulty should not change")
	}
}
