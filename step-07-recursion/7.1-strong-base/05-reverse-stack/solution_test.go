package reversestack

import (
	"slices"
	"testing"
)

func TestReverse(t *testing.T) {
	for _, a := range [][]int{{1, 2, 3, 4}, {}, {7}, {1, 2}} {
		s := Stack(append([]int{}, a...))
		Reverse(&s)
		want := append([]int{}, a...)
		slices.Reverse(want)
		if !slices.Equal([]int(s), want) {
			t.Errorf("Reverse(%v) = %v", a, s)
		}
	}
}
