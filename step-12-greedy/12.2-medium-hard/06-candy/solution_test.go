package candy

import "testing"

func TestCandy(t *testing.T) {
	tests := []struct {
		r    []int
		want int
	}{{[]int{1, 0, 2}, 5}, {[]int{1, 2, 2}, 4}, {[]int{1, 3, 4, 5, 2}, 11}, {[]int{5}, 1}, {[]int{1, 2, 87, 87, 87, 2, 1}, 13}}
	for _, tc := range tests {
		if got := Candy(tc.r); got != tc.want {
			t.Errorf("Candy(%v) = %d", tc.r, got)
		}
	}
}
