package rearrangebysign

import (
	"slices"
	"testing"
)

func TestRearrange(t *testing.T) {
	if got := RearrangeEqual([]int{3, 1, -2, -5, 2, -4}); !slices.Equal(got, []int{3, -2, 1, -5, 2, -4}) {
		t.Errorf("RearrangeEqual = %v", got)
	}
	if got := RearrangeAny([]int{1, 2, -4, -5, 3, 6}); !slices.Equal(got, []int{1, -4, 2, -5, 3, 6}) {
		t.Errorf("RearrangeAny = %v", got)
	}
	if got := RearrangeAny([]int{-1, -2, 3}); !slices.Equal(got, []int{3, -1, -2}) {
		t.Errorf("RearrangeAny more negatives = %v", got)
	}
}
