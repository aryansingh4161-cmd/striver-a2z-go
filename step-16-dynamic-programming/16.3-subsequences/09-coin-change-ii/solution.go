package coinchangeii

// Recursive counts combinations (skip coin i, or use it and stay). Exponential.
func Recursive(coins []int, amount int) int {
	var f func(i, t int) int
	f = func(i, t int) int {
		if i == 0 {
			if t%coins[0] == 0 {
				return 1
			}
			return 0
		}
		res := f(i-1, t)
		if coins[i] <= t {
			res += f(i, t-coins[i])
		}
		return res
	}
	return f(len(coins)-1, amount)
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
				return 1
			}
			return 0
		}
		if dp[i][t] != -1 {
			return dp[i][t]
		}
		res := f(i-1, t)
		if coins[i] <= t {
			res += f(i, t-coins[i])
		}
		dp[i][t] = res
		return res
	}
	return f(len(coins)-1, amount)
}

// Optimal: coins outside, amounts inside going up. O(amount) space.
func Optimal(coins []int, amount int) int {
	ways := make([]int, amount+1)
	ways[0] = 1
	for _, c := range coins {
		for t := c; t <= amount; t++ {
			ways[t] += ways[t-c]
		}
	}
	return ways[amount]
}
