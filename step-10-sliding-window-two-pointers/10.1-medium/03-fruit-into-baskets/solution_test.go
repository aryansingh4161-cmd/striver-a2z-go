package fruitintobaskets

import "testing"

func TestTotalFruit(t *testing.T) {
	tests := []struct {
		f    []int
		want int
	}{{[]int{1, 2, 1}, 3}, {[]int{0, 1, 2, 2}, 3}, {[]int{1, 2, 3, 2, 2}, 4}, {[]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}, 5}, {[]int{}, 0}}
	for _, tc := range tests {
		if got := TotalFruit(tc.f); got != tc.want {
			t.Errorf("TotalFruit(%v) = %d", tc.f, got)
		}
	}
}
