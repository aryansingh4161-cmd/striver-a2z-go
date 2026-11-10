package houserobberii

// linear is the straight-line house robber with O(1) space.
func linear(nums []int) int {
	prev2, prev1 := 0, 0
	for _, x := range nums {
		prev2, prev1 = prev1, max(x+prev2, prev1)
	}
	return prev1
}

// Rob solves the circular version: skip the first or skip the last. O(n).
func Rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return max(linear(nums[:n-1]), linear(nums[1:]))
}
