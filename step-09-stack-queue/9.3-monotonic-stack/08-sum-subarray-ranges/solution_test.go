package sumsubarrayranges

import (
	"math/rand"
	"testing"
)

func TestRanges(t *testing.T) {
	tests := []struct {
		a    []int
		want int
	}{{[]int{1, 2, 3}, 4}, {[]int{1, 3, 3}, 4}, {[]int{4, -2, -3, 4, 1}, 59}}
	for _, tc := range tests {
		if got := Ranges(tc.a); got != tc.want {
			t.Errorf("Ranges(%v) = %d", tc.a, got)
		}
	}
	r := rand.New(rand.NewSource(5))
	for k := 0; k < 200; k++ {
		a := make([]int, r.Intn(20)+1)
		for i := range a {
			a[i] = r.Intn(7) - 3
		}
		if Ranges(a) != RangesBrute(a) {
			t.Fatalf("mismatch on %v", a)
		}
	}
}
