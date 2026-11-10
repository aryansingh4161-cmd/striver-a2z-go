package lis

// Recursive: f(i, prev) with take / skip. O(2^n).
func Recursive(nums []int) int {
	n := len(nums)
	var f func(i, prev int) int
	f = func(i, prev int) int {
		if i == n {
			return 0
		}
		best := f(i+1, prev)
		if prev == -1 || nums[i] > nums[prev] {
			best = max(best, 1+f(i+1, i))
		}
		return best
	}
	return f(0, -1)
}

// Memo caches f(i, prev+1). O(n²).
func Memo(nums []int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(i, prev int) int
	f = func(i, prev int) int {
		if i == n {
			return 0
		}
		if dp[i][prev+1] != -1 {
			return dp[i][prev+1]
		}
		best := f(i+1, prev)
		if prev == -1 || nums[i] > nums[prev] {
			best = max(best, 1+f(i+1, i))
		}
		dp[i][prev+1] = best
		return best
	}
	return f(0, -1)
}

// TwoRows is the bottom-up f(i, prev) keeping only row i+1. O(n²) time, O(n) space.
func TwoRows(nums []int) int {
	n := len(nums)
	ahead := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		cur := make([]int, n+1)
		for prev := i - 1; prev >= -1; prev-- {
			best := ahead[prev+1]
			if prev == -1 || nums[i] > nums[prev] {
				best = max(best, 1+ahead[i+1])
			}
			cur[prev+1] = best
		}
		ahead = cur
	}
	return ahead[0]
}

// EndingAt: dp[i] = LIS ending at i. O(n²) time, O(n) space.
func EndingAt(nums []int) int {
	dp := make([]int, len(nums))
	best := 0
	for i := range nums {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		best = max(best, dp[i])
	}
	return best
}
