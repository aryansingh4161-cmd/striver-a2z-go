package minimumcoins

import (
	"slices"
	"testing"
)

func TestMinCoins(t *testing.T) {
	if got := MinCoins(49); !slices.Equal(got, []int{20, 20, 5, 2, 2}) {
		t.Errorf("49 → %v", got)
	}
	if got := MinCoins(2593); len(got) != 7 { // 2000 500 50 20 20 2 1
		t.Errorf("2593 → %v", got)
	}
	// greedy is not optimal for {4, 3, 1}: 6 = 4+1+1 (3) but 3+3 (2)
	if GreedyCount([]int{4, 3, 1}, 6) != 3 {
		t.Error("expected greedy to use 3 coins")
	}
}
