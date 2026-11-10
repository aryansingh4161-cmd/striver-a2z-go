package numberoflis

// CountLIS returns how many longest strictly increasing subsequences exist. O(n²).
func CountLIS(nums []int) int {
	n := len(nums)
	dp, cnt := make([]int, n), make([]int, n)
	best := 0
	for i := 0; i < n; i++ {
		dp[i], cnt[i] = 1, 1
		for j := 0; j < i; j++ {
			if nums[j] >= nums[i] {
				continue
			}
			switch {
			case dp[j]+1 > dp[i]:
				dp[i], cnt[i] = dp[j]+1, cnt[j]
			case dp[j]+1 == dp[i]:
				cnt[i] += cnt[j]
			}
		}
		best = max(best, dp[i])
	}
	total := 0
	for i := 0; i < n; i++ {
		if dp[i] == best {
			total += cnt[i]
		}
	}
	return total
}
