package firstlastoccurrence

// SearchRange returns the first and last index of target, or [-1, -1].
// Two binary searches. Time O(log n), space O(1).
func SearchRange(a []int, target int) [2]int {
	return [2]int{find(a, target, true), find(a, target, false)}
}

// find returns the first (first=true) or last index of target, or -1.
func find(a []int, target int, first bool) int {
	lo, hi, ans := 0, len(a)-1, -1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch {
		case a[mid] == target:
			ans = mid
			if first {
				hi = mid - 1 // look for an earlier one
			} else {
				lo = mid + 1 // look for a later one
			}
		case a[mid] < target:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return ans
}
