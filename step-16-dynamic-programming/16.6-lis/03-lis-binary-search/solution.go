package lisbinarysearch

import "sort"

// LIS returns the length of the longest strictly increasing subsequence. O(n log n).
func LIS(nums []int) int {
	tails := []int{}
	for _, x := range nums {
		k := sort.SearchInts(tails, x) // first index with tails[k] >= x
		if k == len(tails) {
			tails = append(tails, x)
		} else {
			tails[k] = x
		}
	}
	return len(tails)
}
