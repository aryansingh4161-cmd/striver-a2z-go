package fallingpathsum

const inf = 1 << 60

func pick(wantMax bool, a, b int) int {
	if wantMax {
		return max(a, b)
	}
	return min(a, b)
}

// Recursive tries every start and every move. Exponential.
func Recursive(mat [][]int, wantMax bool) int {
	n, m := len(mat), len(mat[0])
	bad := inf
	if wantMax {
		bad = -inf
	}
	var f func(i, j int) int
	f = func(i, j int) int {
		if j < 0 || j >= m {
			return bad // never chosen
		}
		if i == 0 {
			return mat[0][j]
		}
		return mat[i][j] + pick(wantMax, f(i-1, j), pick(wantMax, f(i-1, j-1), f(i-1, j+1)))
	}
	best := f(n-1, 0)
	for j := 1; j < m; j++ {
		best = pick(wantMax, best, f(n-1, j))
	}
	return best
}

// Tab fills row by row. O(n·m) time and space.
func Tab(mat [][]int, wantMax bool) int {
	n, m := len(mat), len(mat[0])
	dp := make([][]int, n)
	dp[0] = append([]int(nil), mat[0]...)
	for i := 1; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			b := dp[i-1][j]
			if j > 0 {
				b = pick(wantMax, b, dp[i-1][j-1])
			}
			if j+1 < m {
				b = pick(wantMax, b, dp[i-1][j+1])
			}
			dp[i][j] = mat[i][j] + b
		}
	}
	best := dp[n-1][0]
	for _, v := range dp[n-1] {
		best = pick(wantMax, best, v)
	}
	return best
}

// Optimal keeps just the previous row. O(n·m) time, O(m) space.
func Optimal(mat [][]int, wantMax bool) int {
	m := len(mat[0])
	prev := append([]int(nil), mat[0]...)
	cur := make([]int, m)
	for i := 1; i < len(mat); i++ {
		for j := 0; j < m; j++ {
			b := prev[j]
			if j > 0 {
				b = pick(wantMax, b, prev[j-1])
			}
			if j+1 < m {
				b = pick(wantMax, b, prev[j+1])
			}
			cur[j] = mat[i][j] + b
		}
		prev, cur = cur, prev
	}
	best := prev[0]
	for _, v := range prev {
		best = pick(wantMax, best, v)
	}
	return best
}
