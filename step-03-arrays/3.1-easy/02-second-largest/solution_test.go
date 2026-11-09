package secondlargest

import "testing"

func TestSecond(t *testing.T) {
	tests := []struct {
		a            []int
		large, small int
	}{
		{[]int{1, 2, 4, 7, 7, 5}, 5, 2},
		{[]int{10, 10, 10}, -1, -1},
		{[]int{5}, -1, -1},
		{[]int{-3, -1, -2}, -2, -2},
		{[]int{2, 1}, 1, 2},
	}
	for _, tc := range tests {
		if got := SecondLargest(tc.a); got != tc.large {
			t.Errorf("SecondLargest(%v) = %d, want %d", tc.a, got, tc.large)
		}
		if got := SecondSmallest(tc.a); got != tc.small {
			t.Errorf("SecondSmallest(%v) = %d, want %d", tc.a, got, tc.small)
		}
	}
}
