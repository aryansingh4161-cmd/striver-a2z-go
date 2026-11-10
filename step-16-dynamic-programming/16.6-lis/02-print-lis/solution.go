package printlis

// PrintLIS returns one longest strictly increasing subsequence. O(n²).
func PrintLIS(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	dp, parent := make([]int, n), make([]int, n)
	last := 0
	for i := 0; i < n; i++ {
		dp[i], parent[i] = 1, i
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
				dp[i], parent[i] = dp[j]+1, j
			}
		}
		if dp[i] > dp[last] {
			last = i
		}
	}
	out := make([]int, dp[last])
	for k, x := len(out)-1, last; k >= 0; k-- {
		out[k] = nums[x]
		x = parent[x]
	}
	return out
}
