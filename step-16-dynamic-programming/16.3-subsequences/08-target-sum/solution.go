package targetsum

import countsubsetssumk "github.com/aryansingh4161-cmd/striver-a2z-go/step-16-dynamic-programming/16.3-subsequences/04-count-subsets-sum-k"

// Recursive tries + and − for every number. O(2^n).
func Recursive(nums []int, target int) int {
	var f func(i, sum int) int
	f = func(i, sum int) int {
		if i < 0 {
			if sum == target {
				return 1
			}
			return 0
		}
		return f(i-1, sum+nums[i]) + f(i-1, sum-nums[i])
	}
	return f(len(nums)-1, 0)
}

// Ways reduces to counting subsets with sum (total − target) / 2. O(n·total).
func Ways(nums []int, target int) int {
	total := 0
	for _, x := range nums {
		total += x
	}
	if total-target < 0 || (total-target)%2 == 1 {
		return 0
	}
	return countsubsetssumk.Optimal(nums, (total-target)/2)
}
