package kruskalsalgorithm

import (
	"sort"

	disjointset "github.com/aryansingh4161-cmd/striver-a2z-go/step-15-graphs/15.5-mst-dsu/02-disjoint-set"
)

// MST returns the total weight of a minimum spanning tree (forest, if the
// graph is disconnected) of an undirected graph with edges [u, v, w]. O(E log E).
func MST(n int, edges [][3]int) int {
	es := append([][3]int(nil), edges...)
	sort.Slice(es, func(i, j int) bool { return es[i][2] < es[j][2] })
	d := disjointset.New(n)
	sum := 0
	for _, e := range es {
		if d.Union(e[0], e[1]) { // false → would make a cycle
			sum += e[2]
		}
	}
	return sum
}
