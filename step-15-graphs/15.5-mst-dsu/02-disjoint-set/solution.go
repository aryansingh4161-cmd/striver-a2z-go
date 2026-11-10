// Package disjointset is the Union-Find used by the rest of Step 15.5.
package disjointset

// DSU supports Union and Find in amortised ~O(1) (union by size + path compression).
type DSU struct {
	parent, size []int
	Groups       int // number of separate groups right now
}

// New makes n singleton groups 0..n-1.
func New(n int) *DSU {
	d := &DSU{parent: make([]int, n), size: make([]int, n), Groups: n}
	for i := range d.parent {
		d.parent[i], d.size[i] = i, 1
	}
	return d
}

// Find returns the root of x's group, compressing the path on the way.
func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

// Union merges the groups of a and b. It returns false if they were already together.
func (d *DSU) Union(a, b int) bool {
	ra, rb := d.Find(a), d.Find(b)
	if ra == rb {
		return false
	}
	if d.size[ra] < d.size[rb] {
		ra, rb = rb, ra
	}
	d.parent[rb] = ra // smaller tree goes under the bigger one
	d.size[ra] += d.size[rb]
	d.Groups--
	return true
}

// Same reports whether a and b are in the same group.
func (d *DSU) Same(a, b int) bool { return d.Find(a) == d.Find(b) }

// Size returns how many items are in x's group.
func (d *DSU) Size(x int) int { return d.size[d.Find(x)] }

// RankDSU is the "union by rank" version, shown for comparison.
type RankDSU struct{ parent, rank []int }

// NewRank makes n singleton groups.
func NewRank(n int) *RankDSU {
	d := &RankDSU{parent: make([]int, n), rank: make([]int, n)}
	for i := range d.parent {
		d.parent[i] = i
	}
	return d
}

// Find with path compression.
func (d *RankDSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

// Union attaches the lower-rank root under the higher-rank root.
func (d *RankDSU) Union(a, b int) bool {
	ra, rb := d.Find(a), d.Find(b)
	if ra == rb {
		return false
	}
	switch {
	case d.rank[ra] < d.rank[rb]:
		d.parent[ra] = rb
	case d.rank[ra] > d.rank[rb]:
		d.parent[rb] = ra
	default:
		d.parent[rb] = ra
		d.rank[ra]++
	}
	return true
}
