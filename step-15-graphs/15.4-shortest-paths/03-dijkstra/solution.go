package dijkstra

import (
	"container/heap"
	"math"
)

// Edge is a weighted edge to node To.
type Edge struct{ To, W int }

// Inf marks unreachable nodes.
const Inf = math.MaxInt

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

// Dijkstra returns shortest distances from src and the parent of every node
// on its shortest path (-1 for src / unreachable). O((V + E) log V).
func Dijkstra(adj [][]Edge, src int) (dist, parent []int) {
	n := len(adj)
	dist, parent = make([]int, n), make([]int, n)
	for i := range dist {
		dist[i], parent[i] = Inf, -1
	}
	dist[src] = 0
	h := &minHeap{{0, src}}
	for h.Len() > 0 {
		cur := heap.Pop(h).(item)
		if cur.d > dist[cur.u] {
			continue // stale entry
		}
		for _, e := range adj[cur.u] {
			if nd := cur.d + e.W; nd < dist[e.To] {
				dist[e.To], parent[e.To] = nd, cur.u
				heap.Push(h, item{nd, e.To})
			}
		}
	}
	return dist, parent
}

// PathTo rebuilds the path src → target from the parent array (nil if unreachable).
func PathTo(parent []int, src, target int) []int {
	if target != src && parent[target] == -1 {
		return nil
	}
	path := []int{}
	for v := target; v != -1; v = parent[v] {
		path = append(path, v)
	}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

// Undirected builds an adjacency list from [u, v, w] edges.
func Undirected(n int, edges [][3]int) [][]Edge {
	adj := make([][]Edge, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], Edge{e[1], e[2]})
		adj[e[1]] = append(adj[e[1]], Edge{e[0], e[2]})
	}
	return adj
}
