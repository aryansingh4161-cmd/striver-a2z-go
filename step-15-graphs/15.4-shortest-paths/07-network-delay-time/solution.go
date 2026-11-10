package networkdelaytime

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-15-graphs/15.4-shortest-paths/03-dijkstra"

// NetworkDelay returns the time for the signal from k to reach all n nodes
// (labelled 1..n), or -1. Dijkstra, O((V + E) log V).
func NetworkDelay(times [][3]int, n, k int) int {
	adj := make([][]dijkstra.Edge, n)
	for _, t := range times {
		adj[t[0]-1] = append(adj[t[0]-1], dijkstra.Edge{To: t[1] - 1, W: t[2]})
	}
	dist, _ := dijkstra.Dijkstra(adj, k-1)
	worst := 0
	for _, d := range dist {
		if d == dijkstra.Inf {
			return -1
		}
		worst = max(worst, d)
	}
	return worst
}
