package bridges

import (
	"math/rand"
	"slices"
	"testing"
)

func components(n int, edges [][2]int, skip int) int {
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	var f func(int) int
	f = func(x int) int {
		if p[x] != x {
			p[x] = f(p[x])
		}
		return p[x]
	}
	c := n
	for i, e := range edges {
		if i == skip {
			continue
		}
		if a, b := f(e[0]), f(e[1]); a != b {
			p[a] = b
			c--
		}
	}
	return c
}

// brute: an edge is a bridge if removing it adds a component.
func brute(n int, edges [][2]int) [][2]int {
	base := components(n, edges, -1)
	var out [][2]int
	for i, e := range edges {
		if components(n, edges, i) > base {
			a, b := min(e[0], e[1]), max(e[0], e[1])
			out = append(out, [2]int{a, b})
		}
	}
	slices.SortFunc(out, func(x, y [2]int) int {
		if x[0] != y[0] {
			return x[0] - y[0]
		}
		return x[1] - y[1]
	})
	return out
}

func TestBridges(t *testing.T) {
	got := Bridges(4, [][2]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}})
	if !slices.Equal(got, [][2]int{{1, 3}}) {
		t.Errorf("got %v", got)
	}
	if got := Bridges(2, [][2]int{{0, 1}, {0, 1}}); len(got) != 0 {
		t.Errorf("parallel edges are not bridges: %v", got)
	}
	r := rand.New(rand.NewSource(19))
	for it := 0; it < 300; it++ {
		n := 1 + r.Intn(8)
		var edges [][2]int
		for k := r.Intn(12); k > 0; k-- {
			u, v := r.Intn(n), r.Intn(n)
			if u != v {
				edges = append(edges, [2]int{u, v})
			}
		}
		if a, b := Bridges(n, edges), brute(n, edges); !slices.Equal(a, b) {
			t.Fatalf("%v: %v vs %v", edges, a, b)
		}
	}
}
