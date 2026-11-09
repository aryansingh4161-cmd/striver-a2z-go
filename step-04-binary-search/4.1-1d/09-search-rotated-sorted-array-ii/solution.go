package searchrotatedsortedarrayii

// Search reports whether target is in a rotated sorted array that may contain
// duplicates. Time O(log n) average, O(n) worst case.
func Search(a []int, target int) bool {
	lo, hi := 0, len(a)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if a[mid] == target {
			return true
		}
		if a[lo] == a[mid] && a[mid] == a[hi] { // can't tell which half is sorted
			lo++
			hi--
			continue
		}
		if a[lo] <= a[mid] {
			if a[lo] <= target && target < a[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if a[mid] < target && target <= a[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return false
}
