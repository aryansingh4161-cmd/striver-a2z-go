package maximumproductsubarray

import "math"

// MaxProductPrefixSuffix scans prefix and suffix products. Time O(n), space O(1).
func MaxProductPrefixSuffix(a []int) int {
	n := len(a)
	best := math.MinInt
	pre, suf := 1, 1
	for i := 0; i < n; i++ {
		if pre == 0 {
			pre = 1
		}
		if suf == 0 {
			suf = 1
		}
		pre *= a[i]
		suf *= a[n-1-i]
		best = max(best, pre, suf)
	}
	return best
}

// MaxProduct tracks the max and min product ending at each index.
// a must be non-empty. Time O(n), space O(1).
func MaxProduct(a []int) int {
	maxHere, minHere, best := a[0], a[0], a[0]
	for _, x := range a[1:] {
		if x < 0 {
			maxHere, minHere = minHere, maxHere
		}
		maxHere = max(x, maxHere*x)
		minHere = min(x, minHere*x)
		best = max(best, maxHere)
	}
	return best
}
