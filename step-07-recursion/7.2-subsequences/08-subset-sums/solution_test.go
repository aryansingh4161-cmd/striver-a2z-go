package subsetsums

import (
	"slices"
	"testing"
)

func TestSubsetSums(t *testing.T) {
	if got := SubsetSums([]int{2, 3}); !slices.Equal(got, []int{0, 2, 3, 5}) {
		t.Errorf("got %v", got)
	}
	if got := SubsetSums([]int{5, 2, 1}); !slices.Equal(got, []int{0, 1, 2, 3, 5, 6, 7, 8}) {
		t.Errorf("got %v", got)
	}
}
