package moststonesremoved

import disjointset "github.com/aryansingh4161-cmd/striver-a2z-go/step-15-graphs/15.5-mst-dsu/02-disjoint-set"

// RemoveStones returns the most stones that can be removed. Rows and columns
// are DSU nodes; answer = stones − groups. O(n · α).
func RemoveStones(stones [][2]int) int {
	maxR, maxC := 0, 0
	for _, s := range stones {
		maxR, maxC = max(maxR, s[0]), max(maxC, s[1])
	}
	d := disjointset.New(maxR + maxC + 2)
	used := map[int]bool{}
	for _, s := range stones {
		r, c := s[0], maxR+1+s[1]
		d.Union(r, c)
		used[r], used[c] = true, true
	}
	groups := 0
	for node := range used {
		if d.Find(node) == node {
			groups++
		}
	}
	return len(stones) - groups
}
