package nthroot

import "testing"

func TestNthRoot(t *testing.T) {
	tests := []struct{ n, m, want int }{{3, 27, 3}, {4, 69, -1}, {1, 7, 7}, {2, 1, 1}, {10, 1024, 2}, {6, 1000000, 10}}
	for _, tc := range tests {
		if got := NthRoot(tc.n, tc.m); got != tc.want {
			t.Errorf("NthRoot(%d,%d) = %d", tc.n, tc.m, got)
		}
	}
}
