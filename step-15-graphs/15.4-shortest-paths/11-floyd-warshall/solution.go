package floydwarshall

const inf = 1 << 60

// ShortestDistances updates m (−1 = no edge) in place to all-pairs shortest
// distances (−1 = unreachable). It returns true if there's a negative cycle. O(n³).
func ShortestDistances(m [][]int) (negCycle bool) {
	n := len(m)
	for i := range m {
		for j := range m[i] {
			if m[i][j] == -1 && i != j {
				m[i][j] = inf
			}
			if i == j && m[i][j] == -1 {
				m[i][j] = 0
			}
		}
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			if m[i][k] == inf {
				continue
			}
			for j := 0; j < n; j++ {
				if m[k][j] != inf && m[i][k]+m[k][j] < m[i][j] {
					m[i][j] = m[i][k] + m[k][j]
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		if m[i][i] < 0 {
			negCycle = true
		}
	}
	for i := range m {
		for j := range m[i] {
			if m[i][j] == inf {
				m[i][j] = -1
			}
		}
	}
	return negCycle
}
