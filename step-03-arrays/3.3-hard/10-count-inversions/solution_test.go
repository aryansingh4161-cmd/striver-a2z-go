package countinversions

import (
	"math/rand"
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		a    []int
		want int
	}{{[]int{5, 3, 2, 4, 1}, 8}, {[]int{1, 2, 3}, 0}, {[]int{3, 2, 1}, 3}, {[]int{2, 2}, 0}, {[]int{}, 0}}
	for _, tc := range tests {
		if got := Count(tc.a); got != tc.want {
			t.Errorf("Count(%v) = %d", tc.a, got)
		}
	}
	r := rand.New(rand.NewSource(7))
	for k := 0; k < 30; k++ {
		a := make([]int, r.Intn(40))
		for i := range a {
			a[i] = r.Intn(20)
		}
		if Count(a) != CountBrute(a) {
			t.Fatalf("mismatch on %v", a)
		}
	}
}
