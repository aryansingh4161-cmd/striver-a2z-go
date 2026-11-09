package findpeakelement

// PeakLinear returns the first peak index. Time O(n).
func PeakLinear(a []int) int {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return i
		}
	}
	return len(a) - 1
}

// Peak returns the index of some peak using binary search on the slope.
// Time O(log n), space O(1).
func Peak(a []int) int {
	lo, hi := 0, len(a)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if a[mid] < a[mid+1] {
			lo = mid + 1 // uphill: a peak is to the right
		} else {
			hi = mid // downhill: peak at mid or to the left
		}
	}
	return lo
}
