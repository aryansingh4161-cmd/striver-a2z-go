package asteroidcollision

import (
	"slices"
	"testing"
)

func TestCollide(t *testing.T) {
	tests := []struct{ a, want []int }{
		{[]int{5, 10, -5}, []int{5, 10}},
		{[]int{8, -8}, []int{}},
		{[]int{10, 2, -5}, []int{10}},
		{[]int{-2, -1, 1, 2}, []int{-2, -1, 1, 2}},
		{[]int{1, -2, -2, -2}, []int{-2, -2, -2}},
	}
	for _, tc := range tests {
		if got := Collide(tc.a); !slices.Equal(got, tc.want) {
			t.Errorf("Collide(%v) = %v", tc.a, got)
		}
	}
}
