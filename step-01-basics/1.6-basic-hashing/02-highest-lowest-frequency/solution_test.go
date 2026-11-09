package highestlowestfrequency

import "testing"

func TestHighLow(t *testing.T) {
	tests := []struct {
		a           []int
		most, least int
	}{
		{[]int{10, 5, 10, 15, 10, 5}, 10, 15},
		{[]int{2, 2, 3, 4, 4, 2}, 2, 3},
		{[]int{7}, 7, 7},
	}
	for _, tc := range tests {
		m, l := HighLow(tc.a)
		if m != tc.most || l != tc.least {
			t.Errorf("HighLow(%v) = %d,%d want %d,%d", tc.a, m, l, tc.most, tc.least)
		}
	}
}
