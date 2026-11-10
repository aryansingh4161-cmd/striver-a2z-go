package cherrypickupii

const negInf = -(1 << 60)

// Recursive moves both robots together, 9 options a step. O(9^n).
func Recursive(g [][]int) int {
	n, m := len(g), len(g[0])
	var f func(i, a, b int) int
	f = func(i, a, b int) int {
		if a < 0 || a >= m || b < 0 || b >= m {
			return negInf
		}
		v := g[i][a]
		if a != b {
			v += g[i][b]
		}
		if i == n-1 {
			return v
		}
		best := negInf
		for da := -1; da <= 1; da++ {
			for db := -1; db <= 1; db++ {
				best = max(best, f(i+1, a+da, b+db))
			}
		}
		return v + best
	}
	return f(0, 0, m-1)
}

// Memo caches f(i, a, b). O(9·n·m²).
func Memo(g [][]int) int {
	n, m := len(g), len(g[0])
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, m)
		for a := range dp[i] {
			dp[i][a] = make([]int, m)
			for b := range dp[i][a] {
				dp[i][a][b] = -1
			}
		}
	}
	var f func(i, a, b int) int
	f = func(i, a, b int) int {
		if a < 0 || a >= m || b < 0 || b >= m {
			return negInf
		}
		if dp[i][a][b] != -1 {
			return dp[i][a][b]
		}
		v := g[i][a]
		if a != b {
			v += g[i][b]
		}
		if i < n-1 {
			best := negInf
			for da := -1; da <= 1; da++ {
				for db := -1; db <= 1; db++ {
					best = max(best, f(i+1, a+da, b+db))
				}
			}
			v += best
		}
		dp[i][a][b] = v
		return v
	}
	return f(0, 0, m-1)
}

// Optimal is the bottom-up table keeping only the row below: O(m²) space.
func Optimal(g [][]int) int {
	n, m := len(g), len(g[0])
	layer := func() [][]int {
		l := make([][]int, m)
		for a := range l {
			l[a] = make([]int, m)
		}
		return l
	}
	below := layer()
	for a := 0; a < m; a++ {
		for b := 0; b < m; b++ {
			below[a][b] = g[n-1][a]
			if a != b {
				below[a][b] += g[n-1][b]
			}
		}
	}
	for i := n - 2; i >= 0; i-- {
		cur := layer()
		for a := 0; a < m; a++ {
			for b := 0; b < m; b++ {
				best := negInf
				for da := -1; da <= 1; da++ {
					for db := -1; db <= 1; db++ {
						na, nb := a+da, b+db
						if na >= 0 && na < m && nb >= 0 && nb < m {
							best = max(best, below[na][nb])
						}
					}
				}
				v := g[i][a]
				if a != b {
					v += g[i][b]
				}
				cur[a][b] = v + best
			}
		}
		below = cur
	}
	return below[0][m-1]
}
