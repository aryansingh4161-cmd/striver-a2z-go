package leadersinarray

import "math"

// Leaders returns every element greater than all elements to its right,
// in original order. Time O(n), space O(1) extra.
func Leaders(a []int) []int {
	res := []int{}
	maxRight := math.MinInt
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] > maxRight {
			res = append(res, a[i])
			maxRight = a[i]
		}
	}
	// collected right-to-left; reverse to original order
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return res
}
