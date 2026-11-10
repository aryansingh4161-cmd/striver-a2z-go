package sortstack

import (
	"slices"
	"testing"
)

func TestSort(t *testing.T) {
	for _, a := range [][]int{{3, 1, 4, 2}, {}, {5}, {5, 4, 3, 2, 1}, {2, 2, 1}} {
		s := Stack(append([]int{}, a...))
		Sort(&s)
		want := append([]int{}, a...)
		slices.Sort(want)
		if !slices.Equal([]int(s), want) {
			t.Errorf("Sort(%v) = %v", a, s)
		}
	}
}
