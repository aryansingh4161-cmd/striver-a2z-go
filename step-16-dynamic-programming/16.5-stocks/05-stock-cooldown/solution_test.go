package stockcooldown

import (
	"math/rand"
	"testing"
)

func TestCooldown(t *testing.T) {
	fs := []func([]int) int{Recursive, Tab, Optimal}
	cases := []struct {
		p    []int
		want int
	}{{[]int{1, 2, 3, 0, 2}, 3}, {[]int{1}, 0}, {[]int{1, 2, 4}, 3}, {[]int{4, 1, 2, 1, 5}, 4}}
	for _, c := range cases {
		for i, f := range fs {
			if got := f(c.p); got != c.want {
				t.Errorf("f%d(%v) = %d want %d", i, c.p, got, c.want)
			}
		}
	}
	r := rand.New(rand.NewSource(5))
	for it := 0; it < 300; it++ {
		p := make([]int, 1+r.Intn(12))
		for i := range p {
			p[i] = r.Intn(20)
		}
		want := Recursive(p)
		if Tab(p) != want || Optimal(p) != want {
			t.Fatalf("%v", p)
		}
	}
}
