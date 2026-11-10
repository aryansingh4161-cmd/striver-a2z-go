package pathminimumeffort

import (
	"container/heap"
	"math"
)

type item struct{ e, r, c int }
type minHeap []item

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].e < h[j].e }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(item)) }
func (h *minHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

var dirs = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// MinEffort is Dijkstra where a path's cost is its largest step. O(RC log RC).
func MinEffort(h [][]int) int {
	R, C := len(h), len(h[0])
	eff := make([][]int, R)
	for i := range eff {
		eff[i] = make([]int, C)
		for j := range eff[i] {
			eff[i][j] = math.MaxInt
		}
	}
	eff[0][0] = 0
	pq := &minHeap{{0, 0, 0}}
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(item)
		if cur.r == R-1 && cur.c == C-1 {
			return cur.e
		}
		if cur.e > eff[cur.r][cur.c] {
			continue
		}
		for _, d := range dirs {
			r, c := cur.r+d[0], cur.c+d[1]
			if r < 0 || r >= R || c < 0 || c >= C {
				continue
			}
			ne := max(cur.e, abs(h[r][c]-h[cur.r][cur.c]))
			if ne < eff[r][c] {
				eff[r][c] = ne
				heap.Push(pq, item{ne, r, c})
			}
		}
	}
	return 0
}

// MinEffortBinarySearch: smallest k such that BFS using steps ≤ k reaches the end.
func MinEffortBinarySearch(h [][]int) int {
	R, C := len(h), len(h[0])
	can := func(k int) bool {
		seen := make([]bool, R*C)
		seen[0] = true
		q := [][2]int{{0, 0}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			if p[0] == R-1 && p[1] == C-1 {
				return true
			}
			for _, d := range dirs {
				r, c := p[0]+d[0], p[1]+d[1]
				if r < 0 || r >= R || c < 0 || c >= C || seen[r*C+c] || abs(h[r][c]-h[p[0]][p[1]]) > k {
					continue
				}
				seen[r*C+c] = true
				q = append(q, [2]int{r, c})
			}
		}
		return false
	}
	lo, hi := 0, 1_000_000
	for lo < hi {
		mid := (lo + hi) / 2
		if can(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
