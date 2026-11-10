package articulationpoints

import (
	"math/rand"
	"slices"
	"testing"
)

// components among nodes != removed.
func components(n int, edges [][2]int, removed int) int {
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
	if removed >= 0 {
		c--
	}
	for _, e := range edges {
		if e[0] == removed || e[1] == removed {
			continue
		}
		if a, b := f(e[0]), f(e[1]); a != b {
			p[a] = b
			c--
		}
	}
	return c
}

func brute(n int, edges [][2]int) []int {
	base := components(n, edges, -1)
	out := []int{}
	for u := 0; u < n; u++ {
		if components(n, edges, u) > base {
			out = append(out, u)
		}
	}
	return out
}

func TestPoints(t *testing.T) {
	if got := Points(5, [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 1}}); !slices.Equal(got, []int{1}) {
		t.Errorf("got %v", got)
	}
	if got := Points(3, [][2]int{{0, 1}, {1, 2}}); !slices.Equal(got, []int{1}) {
		t.Errorf("path: got %v", got)
	}
	if got := Points(3, [][2]int{{0, 1}, {1, 2}, {2, 0}}); len(got) != 0 {
		t.Errorf("triangle: got %v", got)
	}
	r := rand.New(rand.NewSource(20))
	for it := 0; it < 300; it++ {
		n := 1 + r.Intn(8)
		var edges [][2]int
		for k := r.Intn(12); k > 0; k-- {
			u, v := r.Intn(n), r.Intn(n)
			if u != v {
				edges = append(edges, [2]int{u, v})
			}
		}
		if a, b := Points(n, edges), brute(n, edges); !slices.Equal(a, b) {
			t.Fatalf("%v: %v vs %v", edges, a, b)
		}
	}
}
