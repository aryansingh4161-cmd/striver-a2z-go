package primefactors

import (
	"slices"
	"testing"
)

func TestPrimeFactors(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{{60, []int{2, 3, 5}}, {13, []int{13}}, {1, []int{}}, {1024, []int{2}}, {97 * 89, []int{89, 97}}}
	for _, tc := range tests {
		if got := PrimeFactors(tc.n); !slices.Equal(got, tc.want) {
			t.Errorf("PrimeFactors(%d) = %v", tc.n, got)
		}
	}
}
