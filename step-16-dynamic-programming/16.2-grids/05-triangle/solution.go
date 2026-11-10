package triangle

// Recursive: f(i,j) = tri[i][j] + min(f(i+1,j), f(i+1,j+1)). O(2^n).
func Recursive(tri [][]int) int {
	n := len(tri)
	var f func(i, j int) int
	f = func(i, j int) int {
		if i == n-1 {
			return tri[i][j]
		}
		return tri[i][j] + min(f(i+1, j), f(i+1, j+1))
	}
	return f(0, 0)
}

// Memo caches f(i, j). O(n²).
func Memo(tri [][]int) int {
	n := len(tri)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, i+1)
		for j := range dp[i] {
			dp[i][j] = -1 << 62
		}
	}
	var f func(i, j int) int
	f = func(i, j int) int {
		if i == n-1 {
			return tri[i][j]
		}
		if dp[i][j] != -1<<62 { // values may be negative, so use a sentinel
			return dp[i][j]
		}
		dp[i][j] = tri[i][j] + min(f(i+1, j), f(i+1, j+1))
		return dp[i][j]
	}
	return f(0, 0)
}

// Tab fills from the bottom row up. O(n²) time and space.
func Tab(tri [][]int) int {
	n := len(tri)
	dp := make([][]int, n)
	dp[n-1] = append([]int(nil), tri[n-1]...)
	for i := n - 2; i >= 0; i-- {
		dp[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			dp[i][j] = tri[i][j] + min(dp[i+1][j], dp[i+1][j+1])
		}
	}
	return dp[0][0]
}

// Optimal keeps only the row below. O(n²) time, O(n) space.
func Optimal(tri [][]int) int {
	n := len(tri)
	below := append([]int(nil), tri[n-1]...)
	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			below[j] = tri[i][j] + min(below[j], below[j+1]) // left to right is safe
		}
	}
	return below[0]
}
