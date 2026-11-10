package maxsumnonadjacent

// Recursive: pick (nums[i] + f(i-2)) or skip (f(i-1)). O(2^n).
func Recursive(nums []int) int {
	var f func(int) int
	f = func(i int) int {
		if i < 0 {
			return 0
		}
		if i == 0 {
			return nums[0]
		}
		return max(nums[i]+f(i-2), f(i-1))
	}
	return f(len(nums) - 1)
}

// Memo caches f(i). O(n).
func Memo(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = -1
	}
	var f func(int) int
	f = func(i int) int {
		if i < 0 {
			return 0
		}
		if i == 0 {
			return nums[0]
		}
		if dp[i] != -1 {
			return dp[i]
		}
		dp[i] = max(nums[i]+f(i-2), f(i-1))
		return dp[i]
	}
	return f(len(nums) - 1)
}

// Tab is bottom-up. O(n) time and space.
func Tab(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		pick := nums[i]
		if i >= 2 {
			pick += dp[i-2]
		}
		dp[i] = max(pick, dp[i-1])
	}
	return dp[n-1]
}

// Optimal keeps the last two answers. O(n) time, O(1) space.
func Optimal(nums []int) int {
	prev2, prev1 := 0, nums[0]
	for i := 1; i < len(nums); i++ {
		prev2, prev1 = prev1, max(nums[i]+prev2, prev1)
	}
	return prev1
}
