package minimumcoins

const inf = 1 << 30

// Recursive: skip coin i, or use it and stay at i. Exponential.
func Recursive(coins []int, amount int) int {
	var f func(i, t int) int
	f = func(i, t int) int {
		if i == 0 {
			if t%coins[0] == 0 {
				return t / coins[0]
			}
			return inf
		}
		best := f(i-1, t)
		if coins[i] <= t {
			best = min(best, 1+f(i, t-coins[i]))
		}
		return best
	}
	if r := f(len(coins)-1, amount); r < inf {
		return r
	}
	return -1
}

// Memo caches f(i, t). O(n·amount).
func Memo(coins []int, amount int) int {
	dp := make([][]int, len(coins))
	for i := range dp {
		dp[i] = make([]int, amount+1)
		for t := range dp[i] {
			dp[i][t] = -1
		}
	}
	var f func(i, t int) int
	f = func(i, t int) int {
		if i == 0 {
			if t%coins[0] == 0 {
				return t / coins[0]
			}
			return inf
		}
		if dp[i][t] != -1 {
			return dp[i][t]
		}
		best := f(i-1, t)
		if coins[i] <= t {
			best = min(best, 1+f(i, t-coins[i]))
		}
		dp[i][t] = best
		return best
	}
	if r := f(len(coins)-1, amount); r < inf {
		return r
	}
	return -1
}

// Optimal is the one-row table, t going upward (coins reusable). O(amount) space.
func Optimal(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for t := 1; t <= amount; t++ {
		dp[t] = inf
	}
	for _, c := range coins {
		for t := c; t <= amount; t++ {
			dp[t] = min(dp[t], 1+dp[t-c])
		}
	}
	if dp[amount] >= inf {
		return -1
	}
	return dp[amount]
}
