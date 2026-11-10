package stockiii

import (
	"math/rand"
	"testing"
)

func TestStockIII(t *testing.T) {
	fs := []func([]int) int{Recursive, Optimal, FourStates}
	cases := []struct {
		p    []int
		want int
	}{{[]int{3, 3, 5, 0, 0, 3, 1, 4}, 6}, {[]int{1, 2, 3, 4, 5}, 4}, {[]int{7, 6, 4, 3, 1}, 0}, {[]int{1, 5, 2, 8, 3, 10}, 14}}
	for _, c := range cases {
		for i, f := range fs {
			if got := f(c.p); got != c.want {
				t.Errorf("f%d(%v) = %d want %d", i, c.p, got, c.want)
			}
		}
	}
	r := rand.New(rand.NewSource(3))
	for it := 0; it < 300; it++ {
		p := make([]int, 1+r.Intn(12))
		for i := range p {
			p[i] = r.Intn(20)
		}
		want := Recursive(p)
		if Optimal(p) != want || FourStates(p) != want {
			t.Fatalf("%v", p)
		}
	}
}
