package rodcutting

import (
	"math/rand"
	"testing"
)

func TestRodCutting(t *testing.T) {
	fs := []func([]int, int) int{Recursive, Memo, Optimal}
	cases := []struct {
		p       []int
		n, want int
	}{
		{[]int{2, 5, 7, 8, 10}, 5, 12},
		{[]int{1, 5, 8, 9, 10, 17, 17, 20}, 8, 22},
		{[]int{3}, 1, 3},
	}
	for _, c := range cases {
		for i, f := range fs {
			if got := f(c.p, c.n); got != c.want {
				t.Errorf("f%d = %d want %d", i, got, c.want)
			}
		}
	}
	r := rand.New(rand.NewSource(11))
	for it := 0; it < 200; it++ {
		n := 1 + r.Intn(9)
		p := make([]int, n)
		for i := range p {
			p[i] = r.Intn(30)
		}
		want := Recursive(p, n)
		if Memo(p, n) != want || Optimal(p, n) != want {
			t.Fatal("mismatch")
		}
	}
}
