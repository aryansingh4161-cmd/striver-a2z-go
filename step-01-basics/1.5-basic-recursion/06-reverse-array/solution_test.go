package reversearray

import (
	"slices"
	"testing"
)

func TestReverse(t *testing.T) {
	for _, fn := range []func([]int){ReverseTwoPointers, ReverseOnePointer} {
		a := []int{1, 2, 3, 4, 5}
		fn(a)
		if !slices.Equal(a, []int{5, 4, 3, 2, 1}) {
			t.Errorf("odd length: got %v", a)
		}
		b := []int{1, 2, 3, 4}
		fn(b)
		if !slices.Equal(b, []int{4, 3, 2, 1}) {
			t.Errorf("even length: got %v", b)
		}
		var empty []int
		fn(empty) // must not panic
	}
}
