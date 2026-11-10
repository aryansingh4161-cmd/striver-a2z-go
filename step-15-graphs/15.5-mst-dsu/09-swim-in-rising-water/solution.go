package swiminrisingwater

import (
	"container/heap"

	disjointset "github.com/aryansingh4161-cmd/striver-a2z-go/step-15-graphs/15.5-mst-dsu/02-disjoint-set"
)

var dirs = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// SwimDSU opens cells in height order and unions them with open neighbours;
// the answer is the first time start and end share a group. O(n² · α).
func SwimDSU(grid [][]int) int {
	n := len(grid)
	pos := make([][2]int, n*n) // pos[h] = cell with height h
	for r := range grid {
		for c, h := range grid[r] {
			pos[h] = [2]int{r, c}
		}
	}
	d := disjointset.New(n * n)
	open := make([]bool, n*n)
	for t := 0; t < n*n; t++ {
		r, c := pos[t][0], pos[t][1]
		open[r*n+c] = true
		for _, dd := range dirs {
			nr, nc := r+dd[0], c+dd[1]
			if nr >= 0 && nr < n && nc >= 0 && nc < n && open[nr*n+nc] {
				d.Union(r*n+c, nr*n+nc)
			}
		}
		if open[0] && open[n*n-1] && d.Same(0, n*n-1) {
			return t
		}
	}
	return n*n - 1
}

type item struct{ t, r, c int }
type minHeap []item

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].t < h[j].t }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(item)) }
func (h *minHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// SwimDijkstra: path cost = highest cell on it. O(n² log n).
func SwimDijkstra(grid [][]int) int {
	n := len(grid)
	seen := make([]bool, n*n)
	pq := &minHeap{{grid[0][0], 0, 0}}
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(item)
		if cur.r == n-1 && cur.c == n-1 {
			return cur.t
		}
		if seen[cur.r*n+cur.c] {
			continue
		}
		seen[cur.r*n+cur.c] = true
		for _, dd := range dirs {
			nr, nc := cur.r+dd[0], cur.c+dd[1]
			if nr >= 0 && nr < n && nc >= 0 && nc < n && !seen[nr*n+nc] {
				heap.Push(pq, item{max(cur.t, grid[nr][nc]), nr, nc})
			}
		}
	}
	return -1
}
