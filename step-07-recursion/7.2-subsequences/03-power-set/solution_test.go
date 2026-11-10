package powerset

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

func TestSubsets(t *testing.T) {
	a := []int{1, 2, 3}
	want := canon([][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}})
	if got := canon(SubsetsRecursive(a)); !slices.Equal(got, want) {
		t.Errorf("SubsetsRecursive = %v", got)
	}
	if got := canon(SubsetsBits(a)); !slices.Equal(got, want) {
		t.Errorf("SubsetsBits = %v", got)
	}
	if got := len(SubsetsRecursive([]int{})); got != 1 {
		t.Errorf("empty input should give [[]], got %d sets", got)
	}
}
