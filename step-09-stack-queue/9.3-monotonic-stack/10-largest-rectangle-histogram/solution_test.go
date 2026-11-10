package largestrectanglehistogram

import (
	"math/rand"
	"testing"
)

func TestLargest(t *testing.T) {
	tests := []struct {
		h    []int
		want int
	}{{[]int{2, 1, 5, 6, 2, 3}, 10}, {[]int{2, 4}, 4}, {[]int{}, 0}, {[]int{3, 3, 3}, 9}, {[]int{1, 2, 3, 4, 5}, 9}}
	for _, tc := range tests {
		if got := Largest(tc.h); got != tc.want {
			t.Errorf("Largest(%v) = %d", tc.h, got)
		}
	}
	r := rand.New(rand.NewSource(6))
	for k := 0; k < 200; k++ {
		h := make([]int, r.Intn(20))
		for i := range h {
			h[i] = r.Intn(6)
		}
		if Largest(h) != LargestBrute(h) {
			t.Fatalf("mismatch on %v", h)
		}
	}
}
