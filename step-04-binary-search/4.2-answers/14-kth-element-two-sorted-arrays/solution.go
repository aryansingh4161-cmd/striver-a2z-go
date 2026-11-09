package kthelementtwosortedarrays

import "math"

// KthMerge counts up to k with two pointers. Time O(k).
func KthMerge(a, b []int, k int) int {
	i, j := 0, 0
	for c := 1; ; c++ {
		var v int
		if j >= len(b) || (i < len(a) && a[i] <= b[j]) {
			v = a[i]
			i++
		} else {
			v = b[j]
			j++
		}
		if c == k {
			return v
		}
	}
}

// Kth returns the k-th smallest (1-based) of a and b combined.
// Partition binary search. Time O(log(min(n1,n2))), space O(1).
func Kth(a, b []int, k int) int {
	if len(a) > len(b) {
		a, b = b, a
	}
	n1, n2 := len(a), len(b)
	lo, hi := max(0, k-n2), min(k, n1)
	for lo <= hi {
		cut1 := lo + (hi-lo)/2
		cut2 := k - cut1
		l1, r1 := math.MinInt, math.MaxInt
		l2, r2 := math.MinInt, math.MaxInt
		if cut1 > 0 {
			l1 = a[cut1-1]
		}
		if cut1 < n1 {
			r1 = a[cut1]
		}
		if cut2 > 0 {
			l2 = b[cut2-1]
		}
		if cut2 < n2 {
			r2 = b[cut2]
		}
		switch {
		case l1 > r2:
			hi = cut1 - 1
		case l2 > r1:
			lo = cut1 + 1
		default:
			return max(l1, l2)
		}
	}
	return -1
}
