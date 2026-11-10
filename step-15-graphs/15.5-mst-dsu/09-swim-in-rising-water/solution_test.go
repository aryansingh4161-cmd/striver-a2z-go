package swiminrisingwater

import (
	"math/rand"
	"testing"
)

func TestSwim(t *testing.T) {
	cases := []struct {
		g    [][]int
		want int
	}{
		{[][]int{{0, 2}, {1, 3}}, 3},
		{[][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}}, 16},
		{[][]int{{0}}, 0},
		{[][]int{{3, 2}, {0, 1}}, 3},
	}
	for _, c := range cases {
		if a, b := SwimDSU(c.g), SwimDijkstra(c.g); a != c.want || b != c.want {
			t.Errorf("%v: dsu %d dijkstra %d want %d", c.g, a, b, c.want)
		}
	}
	r := rand.New(rand.NewSource(18))
	for it := 0; it < 200; it++ {
		n := 1 + r.Intn(6)
		perm := r.Perm(n * n)
		g := make([][]int, n)
		for i := range g {
			g[i] = perm[i*n : (i+1)*n]
		}
		if a, b := SwimDSU(g), SwimDijkstra(g); a != b {
			t.Fatalf("%v: %d vs %d", g, a, b)
		}
	}
}
