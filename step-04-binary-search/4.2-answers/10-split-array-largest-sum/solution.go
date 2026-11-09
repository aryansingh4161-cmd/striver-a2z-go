package splitarraylargestsum

// SplitArray returns the minimum possible largest sum when splitting a into
// k contiguous parts. Time O(n log sum), space O(1).
func SplitArray(a []int, k int) int {
	lo, hi := 0, 0
	for _, x := range a {
		lo = max(lo, x)
		hi += x
	}
	ans := hi
	for lo <= hi {
		mid := lo + (hi-lo)/2
		parts, sum := 1, 0
		for _, x := range a {
			if sum+x > mid {
				parts++
				sum = 0
			}
			sum += x
		}
		if parts <= k {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}
