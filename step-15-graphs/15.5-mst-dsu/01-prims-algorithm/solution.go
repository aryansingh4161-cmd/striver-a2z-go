package primsalgorithm

import "container/heap"

type item struct{ w, u, parent int }
type minHeap []item

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].w < h[j].w }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(item)) }
func (h *minHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// MST returns the total weight of a minimum spanning tree of a connected
// undirected graph (edges [u, v, w]) and the chosen edges [parent, child].
// O(E log E).
func MST(n int, edges [][3]int) (int, [][2]int) {
	type edge struct{ to, w int }
	adj := make([][]edge, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], edge{e[1], e[2]})
		adj[e[1]] = append(adj[e[1]], edge{e[0], e[2]})
	}
	inTree := make([]bool, n)
	h := &minHeap{{0, 0, -1}}
	sum, chosen := 0, [][2]int{}
	for h.Len() > 0 {
		cur := heap.Pop(h).(item)
		if inTree[cur.u] {
			continue
		}
		inTree[cur.u] = true
		sum += cur.w
		if cur.parent != -1 {
			chosen = append(chosen, [2]int{cur.parent, cur.u})
		}
		for _, e := range adj[cur.u] {
			if !inTree[e.to] {
				heap.Push(h, item{e.w, e.to, cur.u})
			}
		}
	}
	return sum, chosen
}
