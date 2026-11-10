package combinationsum

import (
	"fmt"
	"slices"
	"sort"
	"testing"
)

func canon(sets [][]int) []string {
	out := make([]string, len(sets))
	for i, s := range sets {
		c := append([]int{}, s...)
		slices.Sort(c)
		out[i] = fmt.Sprint(c)
	}
	sort.Strings(out)
	return out
}

func TestCombinationSum(t *testing.T) {
	if got := canon(CombinationSum([]int{2, 3, 6, 7}, 7)); !slices.Equal(got, canon([][]int{{2, 2, 3}, {7}})) {
		t.Errorf("got %v", got)
	}
	if got := canon(CombinationSum([]int{2, 3, 5}, 8)); !slices.Equal(got, canon([][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}})) {
		t.Errorf("got %v", got)
	}
	if got := CombinationSum([]int{2}, 1); len(got) != 0 {
		t.Errorf("got %v", got)
	}
}
