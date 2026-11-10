package makenetworkconnected

import disjointset "github.com/aryansingh4161-cmd/striver-a2z-go/step-15-graphs/15.5-mst-dsu/02-disjoint-set"

// MakeConnected returns the minimum cable moves to connect all n computers, or -1.
func MakeConnected(n int, connections [][2]int) int {
	if len(connections) < n-1 {
		return -1
	}
	d := disjointset.New(n)
	for _, c := range connections {
		d.Union(c[0], c[1])
	}
	return d.Groups - 1
}
