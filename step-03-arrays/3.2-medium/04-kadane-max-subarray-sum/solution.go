package kadane

import "math"

// MaxSubarrayBrute tries every start with a running sum. Time O(n^2).
func MaxSubarrayBrute(a []int) int {
	best := math.MinInt
	for i := range a {
		sum := 0
		for j := i; j < len(a); j++ {
			sum += a[j]
			best = max(best, sum)
		}
	}
	return best
}

// MaxSubarray is Kadane's algorithm. It returns the best sum and the
// inclusive start/end indices of one best subarray. a must be non-empty.
// Time O(n), space O(1).
func MaxSubarray(a []int) (best, start, end int) {
	best = math.MinInt
	sum, curStart := 0, 0
	for i, x := range a {
		if sum == 0 {
			curStart = i // a new subarray may begin here
		}
		sum += x
		if sum > best {
			best, start, end = sum, curStart, i
		}
		if sum < 0 {
			sum = 0
		}
	}
	return best, start, end
}
