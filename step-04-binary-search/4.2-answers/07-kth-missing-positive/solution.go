package kthmissingpositive

// KthMissingLinear shifts k for every present number <= k. Time O(n).
func KthMissingLinear(a []int, k int) int {
	for _, x := range a {
		if x <= k {
			k++
		} else {
			break
		}
	}
	return k
}

// KthMissing uses binary search on missing-count = a[i] - (i+1).
// Time O(log n), space O(1).
func KthMissing(a []int, k int) int {
	lo, hi := 0, len(a)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if a[mid]-(mid+1) < k {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo + k
}
