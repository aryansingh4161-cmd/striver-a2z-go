package bellmanford

import "math"

// Inf marks unreachable nodes.
const Inf = math.MaxInt

// BellmanFord returns distances from src and whether a negative cycle is
// reachable from src. Edges are directed [u, v, w]. O(V · E).
func BellmanFord(n int, edges [][3]int, src int) (dist []int, negCycle bool) {
	dist = make([]int, n)
	for i := range dist {
		dist[i] = Inf
	}
	dist[src] = 0
	for round := 0; round < n-1; round++ {
		changed := false
		for _, e := range edges {
			u, v, w := e[0], e[1], e[2]
			if dist[u] != Inf && dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
				changed = true
			}
		}
		if !changed {
			break // already final: small speed-up
		}
	}
	for _, e := range edges {
		if dist[e[0]] != Inf && dist[e[0]]+e[2] < dist[e[1]] {
			return dist, true
		}
	}
	return dist, false
}
