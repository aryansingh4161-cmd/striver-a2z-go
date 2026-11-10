package countsubsetssumk

import (
	"math/rand"
	"testing"
)

func brute(a []int, k int) int {
	c := 0
	for mask := 0; mask < 1<<len(a); mask++ {
		s := 0
		for i, x := range a {
			if mask>>i&1 == 1 {
				s += x
			}
		}
		if s == k {
			c++
		}
	}
	return c
}

func TestCount(t *testing.T) {
	fs := []func([]int, int) int{Recursive, Memo, Tab, Optimal}
	cases := []struct {
		a       []int
		k, want int
	}{
		{[]int{1, 2, 2, 3}, 3, 3},
		{[]int{0, 0, 1}, 1, 4},
		{[]int{0, 0, 1}, 0, 4},
		{[]int{1, 1, 1}, 2, 3},
		{[]int{5}, 2, 0},
	}
	for _, c := range cases {
		for i, f := range fs {
			if got := f(c.a, c.k); got != c.want {
				t.Errorf("f%d(%v, %d) = %d want %d", i, c.a, c.k, got, c.want)
			}
		}
	}
	r := rand.New(rand.NewSource(4))
	for it := 0; it < 300; it++ {
		a := make([]int, 1+r.Intn(10))
		for i := range a {
			a[i] = r.Intn(6) // zeros included on purpose
		}
		k := r.Intn(15)
		want := brute(a, k)
		for i, f := range fs {
			if f(a, k) != want {
				t.Fatalf("f%d(%v, %d) want %d", i, a, k, want)
			}
		}
	}
}
