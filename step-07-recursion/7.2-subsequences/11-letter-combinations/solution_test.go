package lettercombinations

import (
	"slices"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	if got := LetterCombinations("23"); !slices.Equal(got, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}) {
		t.Errorf("got %v", got)
	}
	if got := LetterCombinations(""); len(got) != 0 {
		t.Errorf("got %v", got)
	}
	if got := len(LetterCombinations("79")); got != 16 {
		t.Errorf("79 gave %d", got)
	}
}
