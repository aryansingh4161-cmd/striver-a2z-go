package painterspartition

import "testing"

func TestMinTime(t *testing.T) {
	tests := []struct {
		boards  []int
		k, want int
	}{
		{[]int{10, 20, 30, 40}, 2, 60},
		{[]int{5, 5, 5, 5}, 2, 10},
		{[]int{10, 10, 10, 10}, 4, 10},
	}
	for _, tc := range tests {
		if got := MinTime(tc.boards, tc.k); got != tc.want {
			t.Errorf("MinTime(%v,%d) = %d", tc.boards, tc.k, got)
		}
	}
}
