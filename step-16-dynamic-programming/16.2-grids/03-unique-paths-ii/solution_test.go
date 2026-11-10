package uniquepathsii

import (
	"math/rand"
	"testing"
)

func TestUniquePathsII(t *testing.T) {
	fs := []func([][]int) int{Recursive, Memo, Tab, Optimal}
	cases := []struct {
		g    [][]int
		want int
	}{
		{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2},
		{[][]int{{0, 1}, {0, 0}}, 1},
		{[][]int{{1}}, 0},
		{[][]int{{0, 0}, {0, 1}}, 0},
		{[][]int{{0, 1, 0, 0}}, 0},
	}
	for _, c := range cases {
		for i, f := range fs {
			if got := f(c.g); got != c.want {
				t.Errorf("f%d(%v) = %d want %d", i, c.g, got, c.want)
			}
		}
	}
	r := rand.New(rand.NewSource(7))
	for it := 0; it < 300; it++ {
		m, n := 1+r.Intn(6), 1+r.Intn(6)
		g := make([][]int, m)
		for i := range g {
			g[i] = make([]int, n)
			for j := range g[i] {
				if r.Intn(4) == 0 {
					g[i][j] = 1
				}
			}
		}
		want := Recursive(g)
		for i, f := range fs[1:] {
			if f(g) != want {
				t.Fatalf("f%d on %v", i+1, g)
			}
		}
	}
}
