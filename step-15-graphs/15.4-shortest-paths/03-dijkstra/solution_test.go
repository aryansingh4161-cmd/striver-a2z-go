package dijkstra

import (
	"math/rand"
	"slices"
	"testing"
)

func TestDijkstra(t *testing.T) {
	adj := Undirected(6, [][3]int{{0, 1, 4}, {0, 2, 4}, {1, 2, 2}, {2, 3, 3}, {2, 4, 1}, {2, 5, 6}, {3, 5, 2}, {4, 5, 3}})
	dist, parent := Dijkstra(adj, 0)
	if !slices.Equal(dist, []int{0, 4, 4, 7, 5, 8}) {
		t.Errorf("dist %v", dist)
	}
	if p := PathTo(parent, 0, 5); !slices.Equal(p, []int{0, 2, 4, 5}) {
		t.Errorf("path %v", p)
	}
	d, par := Dijkstra(make([][]Edge, 2), 0)
	if d[1] != Inf || PathTo(par, 0, 1) != nil {
		t.Error("unreachable")
	}
}

// Brute force: relax all edges n times (Bellman-Ford style).
func TestRandom(t *testing.T) {
	r := rand.New(rand.NewSource(7))
	for it := 0; it < 300; it++ {
		n := 1 + r.Intn(8)
		var edges [][3]int
		for u := 0; u < n; u++ {
			for v := u + 1; v < n; v++ {
				if r.Intn(2) == 0 {
					edges = append(edges, [3]int{u, v, r.Intn(10)})
				}
			}
		}
		got, _ := Dijkstra(Undirected(n, edges), 0)
		want := make([]int, n)
		for i := range want {
			want[i] = Inf
		}
		want[0] = 0
		for k := 0; k < n; k++ {
			for _, e := range edges {
				for _, p := range [][2]int{{e[0], e[1]}, {e[1], e[0]}} {
					if want[p[0]] != Inf && want[p[0]]+e[2] < want[p[1]] {
						want[p[1]] = want[p[0]] + e[2]
					}
				}
			}
		}
		if !slices.Equal(got, want) {
			t.Fatalf("got %v want %v", got, want)
		}
	}
}
