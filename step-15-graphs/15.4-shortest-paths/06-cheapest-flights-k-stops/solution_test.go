package cheapestflightskstops

import (
	"math/rand"
	"testing"
)

// brute: DFS over all paths with at most k+1 edges.
func brute(n int, fl [][3]int, src, dst, k int) int {
	best := -1
	var dfs func(u, edges, cost int)
	dfs = func(u, edges, cost int) {
		if u == dst && (best == -1 || cost < best) {
			best = cost
		}
		if edges == k+1 {
			return
		}
		for _, f := range fl {
			if f[0] == u {
				dfs(f[1], edges+1, cost+f[2])
			}
		}
	}
	dfs(src, 0, 0)
	return best
}

func TestCheapest(t *testing.T) {
	fl := [][3]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}
	for _, f := range []func(int, [][3]int, int, int, int) int{CheapestBFS, CheapestBellmanFord} {
		if got := f(4, fl, 0, 3, 1); got != 700 {
			t.Errorf("got %d want 700", got)
		}
		if got := f(4, fl, 0, 3, 2); got != 400 {
			t.Errorf("got %d want 400", got)
		}
		if got := f(3, [][3]int{{0, 1, 5}}, 0, 2, 3); got != -1 {
			t.Errorf("got %d want -1", got)
		}
	}
	r := rand.New(rand.NewSource(9))
	for it := 0; it < 300; it++ {
		n := 2 + r.Intn(5)
		var fl [][3]int
		for u := 0; u < n; u++ {
			for v := 0; v < n; v++ {
				if u != v && r.Intn(3) == 0 {
					fl = append(fl, [3]int{u, v, 1 + r.Intn(20)})
				}
			}
		}
		k := r.Intn(n)
		want := brute(n, fl, 0, n-1, k)
		if a, b := CheapestBFS(n, fl, 0, n-1, k), CheapestBellmanFord(n, fl, 0, n-1, k); a != want || b != want {
			t.Fatalf("%v k=%d: bfs %d bf %d want %d", fl, k, a, b, want)
		}
	}
}
