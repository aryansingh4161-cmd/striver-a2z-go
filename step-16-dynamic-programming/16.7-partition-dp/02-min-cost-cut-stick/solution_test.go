package mincostcutstick

import (
	"math/rand"
	"testing"
)

func TestCutStick(t *testing.T) {
	cases := []struct {
		n    int
		cuts []int
		want int
	}{{7, []int{1, 3, 4, 5}, 16}, {9, []int{5, 6, 1, 4, 2}, 22}, {5, nil, 0}}
	for _, c := range cases {
		if Recursive(c.n, c.cuts) != c.want || Tab(c.n, c.cuts) != c.want {
			t.Errorf("n=%d cuts=%v", c.n, c.cuts)
		}
	}
	r := rand.New(rand.NewSource(2))
	for it := 0; it < 200; it++ {
		n := 2 + r.Intn(15)
		perm := r.Perm(n - 1)
		cuts := make([]int, r.Intn(min(6, n-1)+1))
		for i := range cuts {
			cuts[i] = perm[i] + 1
		}
		if Recursive(n, cuts) != Tab(n, cuts) {
			t.Fatalf("n=%d %v", n, cuts)
		}
	}
}
