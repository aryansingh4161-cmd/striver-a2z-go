package frogjumpk

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Recursive tries every jump length 1..k. Exponential.
func Recursive(h []int, k int) int {
	var f func(int) int
	f = func(i int) int {
		if i == 0 {
			return 0
		}
		best := -1
		for j := 1; j <= k && i-j >= 0; j++ {
			c := f(i-j) + abs(h[i]-h[i-j])
			if best == -1 || c < best {
				best = c
			}
		}
		return best
	}
	return f(len(h) - 1)
}

// Memo caches f(i). O(n·k).
func Memo(h []int, k int) int {
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
		best := -1
		for j := 1; j <= k && i-j >= 0; j++ {
			c := f(i-j) + abs(h[i]-h[i-j])
			if best == -1 || c < best {
				best = c
			}
		}
		dp[i] = best
		return best
	}
	return f(len(h) - 1)
}

// Tab is the bottom-up version. O(n·k) time, O(n) space.
func Tab(h []int, k int) int {
	dp := make([]int, len(h))
	for i := 1; i < len(h); i++ {
		dp[i] = dp[i-1] + abs(h[i]-h[i-1])
		for j := 2; j <= k && i-j >= 0; j++ {
			dp[i] = min(dp[i], dp[i-j]+abs(h[i]-h[i-j]))
		}
	}
	return dp[len(h)-1]
}
