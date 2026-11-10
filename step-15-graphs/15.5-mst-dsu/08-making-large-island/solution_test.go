package makinglargeisland

import (
	"math/rand"
	"testing"
)

func biggest(g [][]int) int {
	n := len(g)
	seen := make([][]bool, n)
	for i := range seen {
		seen[i] = make([]bool, n)
	}
	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || r >= n || c < 0 || c >= n || g[r][c] == 0 || seen[r][c] {
			return 0
		}
		seen[r][c] = true
		return 1 + dfs(r+1, c) + dfs(r-1, c) + dfs(r, c+1) + dfs(r, c-1)
	}
	best := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			best = max(best, dfs(i, j))
		}
	}
	return best
}

func brute(g [][]int) int {
	best := biggest(g)
	for i := range g {
		for j := range g {
			if g[i][j] == 0 {
				g[i][j] = 1
				best = max(best, biggest(g))
				g[i][j] = 0
			}
		}
	}
	return best
}

func TestLargestIsland(t *testing.T) {
	cases := []struct {
		g    [][]int
		want int
	}{
		{[][]int{{1, 0}, {0, 1}}, 3},
		{[][]int{{1, 1}, {1, 0}}, 4},
		{[][]int{{1, 1}, {1, 1}}, 4},
		{[][]int{{1, 1, 0}, {0, 1, 0}, {1, 0, 1}}, 6},
		{[][]int{{0}}, 1},
	}
	for _, c := range cases {
		if got := LargestIsland(c.g); got != c.want {
			t.Errorf("%v got %d want %d", c.g, got, c.want)
		}
	}
	r := rand.New(rand.NewSource(17))
	for it := 0; it < 300; it++ {
		n := 1 + r.Intn(5)
		g := make([][]int, n)
		for i := range g {
			g[i] = make([]int, n)
			for j := range g[i] {
				g[i][j] = r.Intn(2)
			}
		}
		if a, b := LargestIsland(g), brute(g); a != b {
			t.Fatalf("%v: %d vs %d", g, a, b)
		}
	}
}
