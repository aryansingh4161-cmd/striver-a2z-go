package makinglargeisland

import disjointset "github.com/aryansingh4161-cmd/striver-a2z-go/step-15-graphs/15.5-mst-dsu/02-disjoint-set"

var dirs = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// LargestIsland returns the largest island possible after flipping at most
// one 0 to 1. O(n² · α).
func LargestIsland(grid [][]int) int {
	n := len(grid)
	d := disjointset.New(n * n)
	in := func(r, c int) bool { return r >= 0 && r < n && c >= 0 && c < n }

	// 1. glue existing islands together
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 0 {
				continue
			}
			for _, dd := range dirs {
				nr, nc := r+dd[0], c+dd[1]
				if in(nr, nc) && grid[nr][nc] == 1 {
					d.Union(r*n+c, nr*n+nc)
				}
			}
		}
	}

	// 2. try every 0
	best, sawZero := 0, false
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 1 {
				continue
			}
			sawZero = true
			roots := map[int]bool{}
			size := 1
			for _, dd := range dirs {
				nr, nc := r+dd[0], c+dd[1]
				if in(nr, nc) && grid[nr][nc] == 1 {
					root := d.Find(nr*n + nc)
					if !roots[root] { // same island from two sides counts once
						roots[root] = true
						size += d.Size(root)
					}
				}
			}
			best = max(best, size)
		}
	}
	if !sawZero {
		return n * n
	}
	return best
}
