package numberofways

import (
	"math/rand"
	"testing"
)

// brute: enumerate all simple paths 0 → n-1, count those of minimum length.
func brute(n int, roads [][3]int) int {
	best, count := -1, 0
	used := make([]bool, n)
	var dfs func(u, d int)
	dfs = func(u, d int) {
		if u == n-1 {
			if best == -1 || d < best {
				best, count = d, 1
			} else if d == best {
				count++
			}
			return
		}
		used[u] = true
		for _, r := range roads {
			for _, p := range [][2]int{{r[0], r[1]}, {r[1], r[0]}} {
				if p[0] == u && !used[p[1]] {
					dfs(p[1], d+r[2])
				}
			}
		}
		used[u] = false
	}
	dfs(0, 0)
	return count
}

func TestCountPaths(t *testing.T) {
	if got := CountPaths(4, [][3]int{{0, 1, 2}, {0, 2, 1}, {2, 1, 1}, {1, 3, 1}, {2, 3, 2}}); got != 3 {
		t.Errorf("got %d", got)
	}
	roads := [][3]int{{0, 6, 7}, {0, 1, 2}, {1, 2, 3}, {1, 3, 3}, {6, 3, 3}, {3, 5, 1}, {6, 5, 1}, {2, 5, 1}, {0, 4, 5}, {4, 6, 2}}
	if got := CountPaths(7, roads); got != 4 {
		t.Errorf("leetcode example: got %d", got)
	}
	if got := CountPaths(1, nil); got != 1 {
		t.Errorf("single node: got %d", got)
	}
	r := rand.New(rand.NewSource(10))
	for it := 0; it < 300; it++ {
		n := 2 + r.Intn(5)
		var roads [][3]int
		for u := 0; u < n; u++ {
			for v := u + 1; v < n; v++ {
				if r.Intn(2) == 0 {
					roads = append(roads, [3]int{u, v, 1 + r.Intn(3)})
				}
			}
		}
		if a, b := CountPaths(n, roads), brute(n, roads); a != b {
			t.Fatalf("%v: got %d want %d", roads, a, b)
		}
	}
}
