package lemonadechange

import "testing"

func TestLemonadeChange(t *testing.T) {
	tests := []struct {
		b    []int
		want bool
	}{{[]int{5, 5, 5, 10, 20}, true}, {[]int{5, 5, 10, 10, 20}, false}, {[]int{10}, false}, {[]int{5, 5, 5, 20}, true}, {[]int{5, 10, 5, 20}, true}}
	for _, tc := range tests {
		if got := LemonadeChange(tc.b); got != tc.want {
			t.Errorf("LemonadeChange(%v) = %v", tc.b, got)
		}
	}
}
