package cheapestflightskstops

import "math"

// CheapestBFS: level-by-level BFS on (stops, node, cost). O(E·k).
func CheapestBFS(n int, flights [][3]int, src, dst, k int) int {
	type edge struct{ to, w int }
	adj := make([][]edge, n)
	for _, f := range flights {
		adj[f[0]] = append(adj[f[0]], edge{f[1], f[2]})
	}
	cost := make([]int, n)
	for i := range cost {
		cost[i] = math.MaxInt
	}
	cost[src] = 0
	type state struct{ stops, node, c int }
	queue := []state{{0, src, 0}}
	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		if s.stops > k {
			continue // already used too many stops
		}
		for _, e := range adj[s.node] {
			if nc := s.c + e.w; nc < cost[e.to] {
				cost[e.to] = nc
				queue = append(queue, state{s.stops + 1, e.to, nc})
			}
		}
	}
	if cost[dst] == math.MaxInt {
		return -1
	}
	return cost[dst]
}

// CheapestBellmanFord runs k+1 rounds, each reading the previous round's prices.
func CheapestBellmanFord(n int, flights [][3]int, src, dst, k int) int {
	price := make([]int, n)
	for i := range price {
		price[i] = math.MaxInt
	}
	price[src] = 0
	for round := 0; round <= k; round++ {
		next := append([]int(nil), price...)
		for _, f := range flights {
			u, v, w := f[0], f[1], f[2]
			if price[u] != math.MaxInt && price[u]+w < next[v] {
				next[v] = price[u] + w
			}
		}
		price = next
	}
	if price[dst] == math.MaxInt {
		return -1
	}
	return price[dst]
}
