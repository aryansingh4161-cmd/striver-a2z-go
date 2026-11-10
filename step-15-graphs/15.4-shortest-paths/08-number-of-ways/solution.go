package numberofways

import (
	"container/heap"
	"math"
)

const mod = 1_000_000_007

type item struct{ d, u int }
type minHeap []item

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].d < h[j].d }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(item)) }
func (h *minHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// CountPaths returns the number of shortest paths from 0 to n-1 (mod 1e9+7).
func CountPaths(n int, roads [][3]int) int {
	type edge struct{ to, w int }
	adj := make([][]edge, n)
	for _, r := range roads {
		adj[r[0]] = append(adj[r[0]], edge{r[1], r[2]})
		adj[r[1]] = append(adj[r[1]], edge{r[0], r[2]})
	}
	dist, ways := make([]int, n), make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[0], ways[0] = 0, 1
	h := &minHeap{{0, 0}}
	for h.Len() > 0 {
		cur := heap.Pop(h).(item)
		if cur.d > dist[cur.u] {
			continue
		}
		for _, e := range adj[cur.u] {
			nd := cur.d + e.w
			switch {
			case nd < dist[e.to]:
				dist[e.to], ways[e.to] = nd, ways[cur.u]
				heap.Push(h, item{nd, e.to})
			case nd == dist[e.to]:
				ways[e.to] = (ways[e.to] + ways[cur.u]) % mod
			}
		}
	}
	return ways[n-1]
}
