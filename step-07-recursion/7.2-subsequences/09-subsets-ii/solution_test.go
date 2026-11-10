package subsetsii

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

func TestSubsetsWithDup(t *testing.T) {
	got := SubsetsWithDup([]int{1, 2, 2})
	want := [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}
	if !slices.Equal(canon(got), canon(want)) {
		t.Errorf("got %v", got)
	}
	if got := SubsetsWithDup([]int{0}); len(got) != 2 {
		t.Errorf("got %v", got)
	}
	if got := SubsetsWithDup([]int{2, 2, 2}); len(got) != 4 {
		t.Errorf("got %v", got)
	}
}
