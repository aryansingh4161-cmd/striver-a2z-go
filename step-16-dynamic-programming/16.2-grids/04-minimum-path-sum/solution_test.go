package minimumpathsum

import (
	"math/rand"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	fs := []func([][]int) int{Recursive, Memo, Tab, Optimal}
	cases := []struct {
		g    [][]int
		want int
	}{
		{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, 12},
		{[][]int{{5}}, 5},
	}
	for _, c := range cases {
		for i, f := range fs {
			if got := f(c.g); got != c.want {
				t.Errorf("f%d = %d want %d", i, got, c.want)
			}
		}
	}
	r := rand.New(rand.NewSource(8))
	for it := 0; it < 300; it++ {
		m, n := 1+r.Intn(6), 1+r.Intn(6)
		g := make([][]int, m)
		for i := range g {
			g[i] = make([]int, n)
			for j := range g[i] {
				g[i][j] = r.Intn(10)
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
