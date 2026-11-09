package longestsubarraysumkpositives

// LongestBrute tries every subarray. Time O(n^2).
func LongestBrute(a []int, k int) int {
	best := 0
	for i := range a {
		sum := 0
		for j := i; j < len(a); j++ {
			sum += a[j]
			if sum == k {
				best = max(best, j-i+1)
			}
		}
	}
	return best
}

// Longest uses a sliding window. Valid only when every element is >= 0.
// Time O(n), space O(1).
func Longest(a []int, k int) int {
	best, sum, left := 0, 0, 0
	for right := range a {
		sum += a[right]
		for sum > k && left <= right {
			sum -= a[left]
			left++
		}
		if sum == k {
			best = max(best, right-left+1)
		}
	}
	return best
}
