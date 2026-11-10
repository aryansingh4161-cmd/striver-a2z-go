package stockii

import (
	"math/rand"
	"testing"
)

func TestStockII(t *testing.T) {
	fs := []func([]int) int{Recursive, Tab, Optimal, Greedy}
	cases := []struct {
		p    []int
		want int
	}{{[]int{7, 1, 5, 3, 6, 4}, 7}, {[]int{1, 2, 3, 4, 5}, 4}, {[]int{7, 6, 4, 3, 1}, 0}}
	for _, c := range cases {
		for i, f := range fs {
			if f(c.p) != c.want {
				t.Errorf("f%d(%v)", i, c.p)
			}
		}
	}
	r := rand.New(rand.NewSource(2))
	for it := 0; it < 300; it++ {
		p := make([]int, 1+r.Intn(12))
		for i := range p {
			p[i] = r.Intn(20)
		}
		want := Recursive(p)
		for i, f := range fs[1:] {
			if f(p) != want {
				t.Fatalf("f%d(%v)", i+1, p)
			}
		}
	}
}
