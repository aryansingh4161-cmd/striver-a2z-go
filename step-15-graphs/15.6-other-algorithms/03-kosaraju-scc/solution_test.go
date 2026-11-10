package kosarajuscc

import (
	"fmt"
	"math/rand"
	"testing"
)

// brute: u and v share an SCC iff each reaches the other.
func brute(adj [][]int) [][]int {
	n := len(adj)
	reach := make([][]bool, n)
	for s := 0; s < n; s++ {
		reach[s] = make([]bool, n)
		stack := []int{s}
		reach[s][s] = true
		for len(stack) > 0 {
			u := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for _, v := range adj[u] {
				if !reach[s][v] {
					reach[s][v] = true
					stack = append(stack, v)
				}
			}
		}
	}
	done := make([]bool, n)
	var out [][]int
	for u := 0; u < n; u++ {
		if done[u] {
			continue
		}
		var c []int
		for v := u; v < n; v++ {
			if reach[u][v] && reach[v][u] {
				c = append(c, v)
				done[v] = true
			}
		}
		out = append(out, c)
	}
	return out
}

func TestSCC(t *testing.T) {
	adj := [][]int{{1}, {2}, {0, 3}, {4}, {5}, {3, 6}, {7}, {6}}
	if got := fmt.Sprint(SCC(adj)); got != "[[0 1 2] [3 4 5] [6 7]]" {
		t.Errorf("got %s", got)
	}
	if Count([][]int{{1}, {2}, {}}) != 3 {
		t.Error("a chain has n SCCs")
	}
	r := rand.New(rand.NewSource(21))
	for it := 0; it < 300; it++ {
		n := 1 + r.Intn(8)
		g := make([][]int, n)
		for u := 0; u < n; u++ {
			for v := 0; v < n; v++ {
				if u != v && r.Intn(4) == 0 {
					g[u] = append(g[u], v)
				}
			}
		}
		if a, b := fmt.Sprint(SCC(g)), fmt.Sprint(brute(g)); a != b {
			t.Fatalf("%v: %s vs %s", g, a, b)
		}
	}
}
