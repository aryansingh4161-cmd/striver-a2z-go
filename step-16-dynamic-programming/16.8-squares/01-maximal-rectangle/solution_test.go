package maximalrectangle

import (
	"math/rand"
	"testing"
)

func TestMaximalRectangle(t *testing.T) {
	mat := [][]int{{1, 0, 1, 0, 0}, {1, 0, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 0, 0, 1, 0}}
	if got := MaximalRectangle(mat); got != 6 {
		t.Errorf("got %d", got)
	}
	if MaximalRectangle([][]int{{0}}) != 0 || MaximalRectangle([][]int{{1}}) != 1 {
		t.Error("single cell")
	}
	if largestInHistogram([]int{2, 1, 5, 6, 2, 3}) != 10 {
		t.Error("histogram")
	}
	r := rand.New(rand.NewSource(7))
	for it := 0; it < 200; it++ {
		n, m := 1+r.Intn(5), 1+r.Intn(5)
		g := make([][]int, n)
		for i := range g {
			g[i] = make([]int, m)
			for j := range g[i] {
				if r.Intn(3) > 0 {
					g[i][j] = 1
				}
			}
		}
		if MaximalRectangle(g) != Brute(g) {
			t.Fatalf("%v", g)
		}
	}
}
