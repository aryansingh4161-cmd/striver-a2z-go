package checksorted

import "testing"

func TestIsSorted(t *testing.T) {
	tests := []struct {
		a    []int
		want bool
	}{{[]int{1, 2, 2, 3}, true}, {[]int{1, 3, 2}, false}, {[]int{}, true}, {[]int{5}, true}}
	for _, tc := range tests {
		if got := IsSorted(tc.a); got != tc.want {
			t.Errorf("IsSorted(%v) = %v", tc.a, got)
		}
	}
	rot := []struct {
		a    []int
		want bool
	}{{[]int{3, 4, 5, 1, 2}, true}, {[]int{2, 1, 3, 4}, false}, {[]int{1, 2, 3}, true}, {[]int{1, 1, 1}, true}}
	for _, tc := range rot {
		if got := IsSortedRotated(tc.a); got != tc.want {
			t.Errorf("IsSortedRotated(%v) = %v", tc.a, got)
		}
	}
}
