package minimumcoins

import (
	"math/rand"
	"testing"
)

func TestMinCoins(t *testing.T) {
	fs := []func([]int, int) int{Recursive, Memo, Optimal}
	cases := []struct {
		c       []int
		a, want int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1, 3, 4}, 6, 2},
		{[]int{7}, 0, 0},
	}
	for _, c := range cases {
		for i, f := range fs {
			if got := f(c.c, c.a); got != c.want {
				t.Errorf("f%d(%v, %d) = %d want %d", i, c.c, c.a, got, c.want)
			}
		}
	}
	r := rand.New(rand.NewSource(7))
	for it := 0; it < 200; it++ {
		c := make([]int, 1+r.Intn(4))
		for i := range c {
			c[i] = 1 + r.Intn(9)
		}
		a := r.Intn(30)
		want := Recursive(c, a)
		if Memo(c, a) != want || Optimal(c, a) != want {
			t.Fatalf("%v %d", c, a)
		}
	}
}
