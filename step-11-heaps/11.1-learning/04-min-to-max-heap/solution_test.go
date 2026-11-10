package mintomaxheap

import (
	"slices"
	"testing"
)

func isMaxHeap(a []int) bool {
	for i := 0; i < len(a)/2; i++ {
		l, r := 2*i+1, 2*i+2
		if a[i] < a[l] || (r < len(a) && a[i] < a[r]) {
			return false
		}
	}
	return true
}

func TestConvert(t *testing.T) {
	a := []int{3, 5, 9, 6, 8, 20, 10, 12, 18, 9}
	orig := append([]int{}, a...)
	ConvertToMaxHeap(a)
	if !isMaxHeap(a) {
		t.Errorf("not a max-heap: %v", a)
	}
	x, y := append([]int{}, a...), orig
	slices.Sort(x)
	slices.Sort(y)
	if !slices.Equal(x, y) {
		t.Error("values changed")
	}
}
