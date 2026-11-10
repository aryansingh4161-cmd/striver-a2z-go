package coinchangeii

import (
	"math/rand"
	"testing"
)

func TestCoinChangeII(t *testing.T) {
	fs := []func([]int, int) int{Recursive, Memo, Optimal}
	cases := []struct {
		c       []int
		a, want int
	}{
		{[]int{1, 2, 5}, 5, 4},
		{[]int{2}, 3, 0},
		{[]int{10}, 10, 1},
		{[]int{3, 5}, 0, 1},
	}
	for _, c := range cases {
		for i, f := range fs {
			if got := f(c.c, c.a); got != c.want {
				t.Errorf("f%d(%v,%d) = %d want %d", i, c.c, c.a, got, c.want)
			}
		}
	}
	r := rand.New(rand.NewSource(9))
	for it := 0; it < 200; it++ {
		c := make([]int, 1+r.Intn(4))
		for i := range c {
			c[i] = 1 + r.Intn(9)
		}
		a := r.Intn(25)
		want := Recursive(c, a)
		if Memo(c, a) != want || Optimal(c, a) != want {
			t.Fatalf("%v %d", c, a)
		}
	}
}
