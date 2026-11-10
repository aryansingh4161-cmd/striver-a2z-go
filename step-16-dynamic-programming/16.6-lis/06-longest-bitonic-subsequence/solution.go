package longestbitonicsubsequence

// LongestBitonic = max over peaks of inc[i] + dec[i] − 1. O(n²).
func LongestBitonic(nums []int) int {
	n := len(nums)
	inc, dec := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		inc[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				inc[i] = max(inc[i], inc[j]+1)
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		dec[i] = 1
		for j := n - 1; j > i; j-- {
			if nums[j] < nums[i] {
				dec[i] = max(dec[i], dec[j]+1)
			}
		}
	}
	best := 0
	for i := 0; i < n; i++ {
		best = max(best, inc[i]+dec[i]-1)
	}
	return best
}
