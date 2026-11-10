package frogjump

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Recursive tries both jumps everywhere. O(2^n).
func Recursive(h []int) int {
	var f func(int) int
	f = func(i int) int {
		if i == 0 {
			return 0
		}
		best := f(i-1) + abs(h[i]-h[i-1])
		if i >= 2 {
			best = min(best, f(i-2)+abs(h[i]-h[i-2]))
		}
		return best
	}
	return f(len(h) - 1)
}

// Memo caches f(i). O(n).
func Memo(h []int) int {
	dp := make([]int, len(h))
	for i := range dp {
		dp[i] = -1
	}
	var f func(int) int
	f = func(i int) int {
		if i == 0 {
			return 0
		}
		if dp[i] != -1 {
			return dp[i]
		}
		best := f(i-1) + abs(h[i]-h[i-1])
		if i >= 2 {
			best = min(best, f(i-2)+abs(h[i]-h[i-2]))
		}
		dp[i] = best
		return best
	}
	return f(len(h) - 1)
}

// Tab fills dp from stone 0 upward. O(n) time and space.
func Tab(h []int) int {
	dp := make([]int, len(h))
	for i := 1; i < len(h); i++ {
		dp[i] = dp[i-1] + abs(h[i]-h[i-1])
		if i >= 2 {
			dp[i] = min(dp[i], dp[i-2]+abs(h[i]-h[i-2]))
		}
	}
	return dp[len(h)-1]
}

// Optimal keeps only the last two dp values. O(n) time, O(1) space.
func Optimal(h []int) int {
	prev2, prev1 := 0, 0
	for i := 1; i < len(h); i++ {
		cur := prev1 + abs(h[i]-h[i-1])
		if i >= 2 {
			cur = min(cur, prev2+abs(h[i]-h[i-2]))
		}
		prev2, prev1 = prev1, cur
	}
	return prev1
}
