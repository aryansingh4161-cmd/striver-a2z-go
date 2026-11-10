package countsquares

import (
	"math/rand"
	"testing"
)

func TestCountSquares(t *testing.T) {
	if got := CountSquares([][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}}); got != 15 {
		t.Errorf("got %d", got)
	}
	if got := CountSquares([][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}); got != 7 {
		t.Errorf("got %d", got)
	}
	r := rand.New(rand.NewSource(8))
	for it := 0; it < 200; it++ {
		n, m := 1+r.Intn(6), 1+r.Intn(6)
		g := make([][]int, n)
		for i := range g {
			g[i] = make([]int, m)
			for j := range g[i] {
				if r.Intn(4) > 0 {
					g[i][j] = 1
				}
			}
		}
		if CountSquares(g) != Brute(g) {
			t.Fatalf("%v", g)
		}
	}
}
