package maximumproductsubarray

import "testing"

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		a    []int
		want int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
		{[]int{-2, 3, -4}, 24},
		{[]int{-2}, -2},
		{[]int{0, 2}, 2},
		{[]int{-1, -3, -10, 0, 60}, 60},
	}
	for _, tc := range tests {
		if got := MaxProduct(tc.a); got != tc.want {
			t.Errorf("MaxProduct(%v) = %d", tc.a, got)
		}
		if got := MaxProductPrefixSuffix(tc.a); got != tc.want {
			t.Errorf("MaxProductPrefixSuffix(%v) = %d", tc.a, got)
		}
	}
}
