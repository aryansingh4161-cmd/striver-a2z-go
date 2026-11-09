package nthroot

// NthRoot returns x such that x^n == m, or -1. Time O(n log m).
func NthRoot(n, m int) int {
	lo, hi := 1, m
	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch compare(mid, n, m) {
		case 0:
			return mid
		case -1:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return -1
}

// compare returns -1, 0 or 1 for x^n < m, == m, > m without overflowing.
func compare(x, n, m int) int {
	p := 1
	for i := 0; i < n; i++ {
		p *= x
		if p > m {
			return 1
		}
	}
	if p == m {
		return 0
	}
	return -1
}
