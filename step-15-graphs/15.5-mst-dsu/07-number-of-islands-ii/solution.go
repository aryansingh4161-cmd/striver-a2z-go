package numberofislandsii

import disjointset "github.com/aryansingh4161-cmd/striver-a2z-go/step-15-graphs/15.5-mst-dsu/02-disjoint-set"

// NumIslands returns the island count after each land-adding operation.
// O(k · α) for k operations.
func NumIslands(n, m int, ops [][2]int) []int {
	d := disjointset.New(n * m)
	land := make([]bool, n*m)
	islands := 0
	out := make([]int, 0, len(ops))
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, op := range ops {
		r, c := op[0], op[1]
		id := r*m + c
		if land[id] {
			out = append(out, islands) // added twice: nothing changes
			continue
		}
		land[id] = true
		islands++
		for _, dd := range dirs {
			nr, nc := r+dd[0], c+dd[1]
			if nr < 0 || nr >= n || nc < 0 || nc >= m || !land[nr*m+nc] {
				continue
			}
			if d.Union(id, nr*m+nc) {
				islands--
			}
		}
		out = append(out, islands)
	}
	return out
}
