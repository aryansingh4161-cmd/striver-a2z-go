package mediantwosortedarrays

import "math"

// MedianMerge walks both arrays without storing them. Time O(n1+n2), space O(1).
func MedianMerge(a, b []int) float64 {
	n := len(a) + len(b)
	i1, i2 := (n-1)/2, n/2 // the two middle positions (equal when n is odd)
	var v1, v2 int
	i, j := 0, 0
	for k := 0; k <= i2; k++ {
		var v int
		if j >= len(b) || (i < len(a) && a[i] <= b[j]) {
			v = a[i]
			i++
		} else {
			v = b[j]
			j++
		}
		if k == i1 {
			v1 = v
		}
		if k == i2 {
			v2 = v
		}
	}
	return float64(v1+v2) / 2
}

// Median uses binary search on the partition of the smaller array.
// Time O(log(min(n1,n2))), space O(1).
func Median(a, b []int) float64 {
	if len(a) > len(b) {
		a, b = b, a // search on the smaller array
	}
	n1, n2 := len(a), len(b)
	half := (n1 + n2 + 1) / 2
	lo, hi := 0, n1
	for lo <= hi {
		cut1 := lo + (hi-lo)/2
		cut2 := half - cut1
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
			if (n1+n2)%2 == 1 {
				return float64(max(l1, l2))
			}
			return float64(max(l1, l2)+min(r1, r2)) / 2
		}
	}
	return 0
}
