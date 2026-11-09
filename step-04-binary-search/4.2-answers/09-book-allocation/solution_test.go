package bookallocation

import "testing"

func TestMinMaxPages(t *testing.T) {
	tests := []struct {
		pages   []int
		m, want int
	}{
		{[]int{25, 46, 28, 49, 24}, 4, 71},
		{[]int{12, 34, 67, 90}, 2, 113},
		{[]int{15, 17, 20}, 2, 32},
		{[]int{1, 2}, 3, -1},
	}
	for _, tc := range tests {
		if got := MinMaxPages(tc.pages, tc.m); got != tc.want {
			t.Errorf("MinMaxPages(%v,%d) = %d", tc.pages, tc.m, got)
		}
	}
}
