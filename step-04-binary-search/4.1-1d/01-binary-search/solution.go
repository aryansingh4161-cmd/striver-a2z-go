package binarysearch

// Search returns the index of target in sorted a, or -1. Time O(log n), space O(1).
func Search(a []int, target int) int {
	lo, hi := 0, len(a)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch {
		case a[mid] == target:
			return mid
		case a[mid] < target:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return -1
}

// SearchRecursive is the recursive version. Space O(log n) stack.
func SearchRecursive(a []int, target int) int {
	var f func(lo, hi int) int
	f = func(lo, hi int) int {
		if lo > hi {
			return -1
		}
		mid := lo + (hi-lo)/2
		if a[mid] == target {
			return mid
		}
		if a[mid] < target {
			return f(mid+1, hi)
		}
		return f(lo, mid-1)
	}
	return f(0, len(a)-1)
}
