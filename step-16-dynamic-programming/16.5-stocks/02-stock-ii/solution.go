package stockii

// Recursive: state (day, canBuy). O(2^n).
func Recursive(p []int) int {
	var f func(i int, canBuy bool) int
	f = func(i int, canBuy bool) int {
		if i == len(p) {
			return 0
		}
		if canBuy {
			return max(-p[i]+f(i+1, false), f(i+1, true))
		}
		return max(p[i]+f(i+1, true), f(i+1, false))
	}
	return f(0, true)
}

// Tab: dp[i][1] = can buy, dp[i][0] = holding. O(n) time and space.
func Tab(p []int) int {
	n := len(p)
	dp := make([][2]int, n+1)
	for i := n - 1; i >= 0; i-- {
		dp[i][1] = max(-p[i]+dp[i+1][0], dp[i+1][1])
		dp[i][0] = max(p[i]+dp[i+1][1], dp[i+1][0])
	}
	return dp[0][1]
}

// Optimal keeps only "tomorrow". O(n) time, O(1) space.
func Optimal(p []int) int {
	aheadBuy, aheadSell := 0, 0
	for i := len(p) - 1; i >= 0; i-- {
		buy := max(-p[i]+aheadSell, aheadBuy)
		sell := max(p[i]+aheadBuy, aheadSell)
		aheadBuy, aheadSell = buy, sell
	}
	return aheadBuy
}

// Greedy adds every upward step (valid only for unlimited transactions).
func Greedy(p []int) int {
	profit := 0
	for i := 1; i < len(p); i++ {
		profit += max(0, p[i]-p[i-1])
	}
	return profit
}
