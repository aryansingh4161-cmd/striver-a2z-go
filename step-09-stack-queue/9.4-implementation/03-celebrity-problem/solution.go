package celebrityproblem

// CelebrityCount uses knows / known-by counters. Time O(n^2), space O(n).
func CelebrityCount(m [][]int) int {
	n := len(m)
	knows, knownBy := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if m[i][j] == 1 {
				knows[i]++
				knownBy[j]++
			}
		}
	}
	for i := 0; i < n; i++ {
		if knows[i] == 0 && knownBy[i] == n-1 {
			return i
		}
	}
	return -1
}

// Celebrity eliminates one person per question, then verifies. Time O(n), space O(1).
func Celebrity(m [][]int) int {
	n := len(m)
	a, b := 0, n-1
	for a < b {
		if m[a][b] == 1 {
			a++ // a knows someone → not a celebrity
		} else {
			b-- // b is not known by a → not a celebrity
		}
	}
	for i := 0; i < n; i++ {
		if i != a && (m[a][i] == 1 || m[i][a] == 0) {
			return -1
		}
	}
	return a
}
