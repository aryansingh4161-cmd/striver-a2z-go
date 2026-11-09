package mininrotatedsortedarray

import "math"

// FindMin returns the minimum of a rotated sorted array (distinct values).
// Time O(log n), space O(1).
func FindMin(a []int) int {
	lo, hi, ans := 0, len(a)-1, math.MaxInt
	for lo <= hi {
		if a[lo] <= a[hi] { // whole range sorted
			ans = min(ans, a[lo])
			break
		}
		mid := lo + (hi-lo)/2
		if a[lo] <= a[mid] { // left half sorted
			ans = min(ans, a[lo])
			lo = mid + 1
		} else { // right half sorted
			ans = min(ans, a[mid])
			hi = mid - 1
		}
	}
	return ans
}
