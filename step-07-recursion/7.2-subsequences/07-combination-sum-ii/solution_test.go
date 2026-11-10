package combinationsumii

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

func TestCombinationSum2(t *testing.T) {
	got := CombinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	if !slices.Equal(canon(got), canon([][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}})) {
		t.Errorf("got %v", got)
	}
	got = CombinationSum2([]int{2, 5, 2, 1, 2}, 5)
	if !slices.Equal(canon(got), canon([][]int{{1, 2, 2}, {5}})) {
		t.Errorf("got %v", got)
	}
	if len(CombinationSum2([]int{3}, 2)) != 0 {
		t.Error("expected none")
	}
}
