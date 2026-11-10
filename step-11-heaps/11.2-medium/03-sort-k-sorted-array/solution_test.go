package sortksortedarray

import (
	"slices"
	"testing"
)

func TestSortKSorted(t *testing.T) {
	if got := SortKSorted([]int{6, 5, 3, 2, 8, 10, 9}, 3); !slices.Equal(got, []int{2, 3, 5, 6, 8, 9, 10}) {
		t.Errorf("got %v", got)
	}
	if got := SortKSorted([]int{2, 1, 4, 3}, 1); !slices.Equal(got, []int{1, 2, 3, 4}) {
		t.Errorf("got %v", got)
	}
	if got := SortKSorted(nil, 2); len(got) != 0 {
		t.Errorf("got %v", got)
	}
}
