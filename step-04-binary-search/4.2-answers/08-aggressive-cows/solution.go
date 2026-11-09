package aggressivecows

import "sort"

// MaxMinDistance returns the largest possible minimum distance between k cows.
// Time O(n log n + n log range), space O(n) for the sorted copy.
func MaxMinDistance(stalls []int, k int) int {
	s := append([]int{}, stalls...)
	sort.Ints(s)
	lo, hi, ans := 1, s[len(s)-1]-s[0], 0
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if canPlace(s, mid, k) {
			ans = mid // works; try a bigger distance
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return ans
}

func canPlace(s []int, dist, k int) bool {
	count, last := 1, s[0]
	for _, x := range s[1:] {
		if x-last >= dist {
			count++
			last = x
			if count >= k {
				return true
			}
		}
	}
	return count >= k
}
