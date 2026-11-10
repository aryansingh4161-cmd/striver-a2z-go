package assigncookies

import "testing"

func TestFindContentChildren(t *testing.T) {
	tests := []struct {
		g, s []int
		want int
	}{{[]int{1, 2, 3}, []int{1, 1}, 1}, {[]int{1, 2}, []int{1, 2, 3}, 2}, {[]int{10, 9, 8, 7}, []int{5, 6, 7, 8}, 2}, {[]int{1}, []int{}, 0}}
	for _, tc := range tests {
		if got := FindContentChildren(tc.g, tc.s); got != tc.want {
			t.Errorf("FindContentChildren(%v,%v) = %d", tc.g, tc.s, got)
		}
	}
}
