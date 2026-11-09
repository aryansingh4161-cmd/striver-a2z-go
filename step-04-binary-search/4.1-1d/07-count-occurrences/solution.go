package countoccurrences

// Count returns how many times x appears in sorted a.
// upperBound - lowerBound. Time O(log n).
func Count(a []int, x int) int {
	return bound(a, func(v int) bool { return v > x }) - bound(a, func(v int) bool { return v >= x })
}

// bound returns the first index where ok(a[i]) is true, or len(a).
// ok must be false...false true...true over the sorted slice.
func bound(a []int, ok func(int) bool) int {
	lo, hi, ans := 0, len(a)-1, len(a)
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if ok(a[mid]) {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}
