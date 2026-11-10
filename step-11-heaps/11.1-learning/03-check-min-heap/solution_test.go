package checkminheap

import "testing"

func TestIsMinHeap(t *testing.T) {
	tests := []struct {
		a    []int
		want bool
	}{{[]int{10, 20, 30, 21, 23}, true}, {[]int{10, 20, 30, 25, 15}, false}, {[]int{}, true}, {[]int{1}, true}, {[]int{2, 1}, false}, {[]int{1, 1, 1}, true}}
	for _, tc := range tests {
		if got := IsMinHeap(tc.a); got != tc.want {
			t.Errorf("IsMinHeap(%v) = %v", tc.a, got)
		}
	}
}
