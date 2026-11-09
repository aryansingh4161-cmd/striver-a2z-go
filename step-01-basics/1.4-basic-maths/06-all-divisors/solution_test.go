package alldivisors

import (
	"slices"
	"testing"
)

func TestDivisors(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{36, []int{1, 2, 3, 4, 6, 9, 12, 18, 36}},
		{7, []int{1, 7}},
		{1, []int{1}},
		{12, []int{1, 2, 3, 4, 6, 12}},
	}
	for _, tc := range tests {
		if got := Divisors(tc.n); !slices.Equal(got, tc.want) {
			t.Errorf("Divisors(%d) = %v, want %v", tc.n, got, tc.want)
		}
		if got := DivisorsBrute(tc.n); !slices.Equal(got, tc.want) {
			t.Errorf("DivisorsBrute(%d) = %v, want %v", tc.n, got, tc.want)
		}
	}
}
