package graphrepresentation

import (
	"slices"
	"testing"
)

func TestRepresentations(t *testing.T) {
	edges := [][2]int{{0, 1}, {0, 2}, {1, 2}, {2, 3}}
	m := AdjMatrix(4, edges)
	if m[0][1] != 1 || m[1][0] != 1 || m[0][3] != 0 {
		t.Error("matrix")
	}
	adj := AdjList(4, edges, false)
	if !slices.Equal(adj[2], []int{0, 1, 3}) {
		t.Errorf("list %v", adj[2])
	}
	total := 0
	for _, a := range adj {
		total += len(a)
	}
	if total != 2*len(edges) {
		t.Error("sum of degrees must be 2m")
	}
	dadj := AdjList(4, edges, true)
	if !slices.Equal(InDegrees(dadj), []int{0, 1, 2, 1}) {
		t.Errorf("in-degrees %v", InDegrees(dadj))
	}
	w := WeightedAdj(2, [][3]int{{0, 1, 5}}, false)
	if w[1][0] != (Edge{0, 5}) {
		t.Error("weighted")
	}
}
