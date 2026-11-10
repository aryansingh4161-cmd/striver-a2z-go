package cherrypickupii

import (
	"math/rand"
	"testing"
)

func TestCherry(t *testing.T) {
	fs := []func([][]int) int{Recursive, Memo, Optimal}
	cases := []struct {
		g    [][]int
		want int
	}{
		{[][]int{{3, 1, 1}, {2, 5, 1}, {1, 5, 5}, {2, 1, 1}}, 24},
		{[][]int{{1, 0, 0, 0, 0, 0, 1}, {2, 0, 0, 0, 0, 3, 0}, {2, 0, 9, 0, 0, 0, 0}, {0, 3, 0, 5, 4, 0, 0}, {1, 0, 2, 3, 0, 0, 6}}, 28},
		{[][]int{{7}}, 7}, // both on the same cell: counted once
		{[][]int{{2, 3}}, 5},
	}
	for _, c := range cases {
		for i, f := range fs {
			if got := f(c.g); got != c.want {
				t.Errorf("f%d = %d want %d", i, got, c.want)
			}
		}
	}
	r := rand.New(rand.NewSource(11))
	for it := 0; it < 100; it++ {
		n, m := 1+r.Intn(5), 1+r.Intn(5)
		g := make([][]int, n)
		for i := range g {
			g[i] = make([]int, m)
			for j := range g[i] {
				g[i][j] = r.Intn(10)
			}
		}
		want := Recursive(g)
		if Memo(g) != want || Optimal(g) != want {
			t.Fatalf("mismatch on %v", g)
		}
	}
}
