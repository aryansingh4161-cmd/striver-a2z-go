package mincostcutstick

import "sort"

const inf = 1 << 60

func withEnds(n int, cuts []int) []int {
	c := append([]int{0}, cuts...)
	sort.Ints(c)
	return append(c, n)
}

// Recursive chooses the first cut in every range. Exponential.
func Recursive(n int, cuts []int) int {
	c := withEnds(n, cuts)
	var f func(i, j int) int
	f = func(i, j int) int {
		if i > j {
			return 0
		}
		best := inf
		for k := i; k <= j; k++ {
			best = min(best, c[j+1]-c[i-1]+f(i, k-1)+f(k+1, j))
		}
		return best
	}
	return f(1, len(cuts))
}

// Tab fills ranges from the inside out. O(m³).
func Tab(n int, cuts []int) int {
	c := withEnds(n, cuts)
	m := len(cuts)
	dp := make([][]int, m+2)
	for i := range dp {
		dp[i] = make([]int, m+2)
	}
	for i := m; i >= 1; i-- {
		for j := i; j <= m; j++ {
			best := inf
			for k := i; k <= j; k++ {
				best = min(best, c[j+1]-c[i-1]+dp[i][k-1]+dp[k+1][j])
			}
			dp[i][j] = best
		}
	}
	return dp[1][m]
}
