package numberofislandsii

import (
	"math/rand"
	"slices"
	"testing"
)

// brute: recount with DFS after every operation.
func brute(n, m int, ops [][2]int) []int {
	g := make([][]bool, n)
	for i := range g {
		g[i] = make([]bool, m)
	}
	var out []int
	for _, op := range ops {
		g[op[0]][op[1]] = true
		seen := make([][]bool, n)
		for i := range seen {
			seen[i] = make([]bool, m)
		}
		var dfs func(r, c int)
		dfs = func(r, c int) {
			if r < 0 || r >= n || c < 0 || c >= m || !g[r][c] || seen[r][c] {
				return
			}
			seen[r][c] = true
			dfs(r+1, c)
			dfs(r-1, c)
			dfs(r, c+1)
			dfs(r, c-1)
		}
		cnt := 0
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if g[i][j] && !seen[i][j] {
					cnt++
					dfs(i, j)
				}
			}
		}
		out = append(out, cnt)
	}
	return out
}

func TestNumIslands(t *testing.T) {
	if got := NumIslands(4, 5, [][2]int{{1, 1}, {0, 1}, {3, 3}, {3, 4}}); !slices.Equal(got, []int{1, 1, 2, 2}) {
		t.Errorf("got %v", got)
	}
	r := rand.New(rand.NewSource(16))
	for it := 0; it < 200; it++ {
		n, m := 1+r.Intn(5), 1+r.Intn(5)
		ops := make([][2]int, 1+r.Intn(15))
		for i := range ops {
			ops[i] = [2]int{r.Intn(n), r.Intn(m)}
		}
		if a, b := NumIslands(n, m, ops), brute(n, m, ops); !slices.Equal(a, b) {
			t.Fatalf("%v: %v vs %v", ops, a, b)
		}
	}
}
