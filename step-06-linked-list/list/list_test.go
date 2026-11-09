package list

import (
	"slices"
	"testing"
)

func TestHelpers(t *testing.T) {
	if got := ToSlice(FromSlice([]int{1, 2, 3})); !slices.Equal(got, []int{1, 2, 3}) {
		t.Errorf("singly = %v", got)
	}
	if FromSlice(nil) != nil {
		t.Error("empty should be nil")
	}
	vals, ok := DToSlice(DFromSlice([]int{4, 5, 6}))
	if !ok || !slices.Equal(vals, []int{4, 5, 6}) {
		t.Errorf("doubly = %v ok=%v", vals, ok)
	}
}
