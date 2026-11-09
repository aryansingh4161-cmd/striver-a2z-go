package lowerbound

// LowerBound returns the first index i with a[i] >= x, or len(a).
// Time O(log n), space O(1).
func LowerBound(a []int, x int) int {
	lo, hi, ans := 0, len(a)-1, len(a)
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if a[mid] >= x {
			ans = mid // possible answer; look for an earlier one
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}
