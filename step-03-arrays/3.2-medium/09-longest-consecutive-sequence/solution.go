package longestconsecutivesequence

import "sort"

// LongestSort sorts a copy and counts runs. Time O(n log n).
func LongestSort(a []int) int {
	if len(a) == 0 {
		return 0
	}
	b := append([]int{}, a...)
	sort.Ints(b)
	best, count := 1, 1
	for i := 1; i < len(b); i++ {
		switch {
		case b[i] == b[i-1]: // duplicate, ignore
		case b[i] == b[i-1]+1:
			count++
		default:
			count = 1
		}
		best = max(best, count)
	}
	return best
}

// Longest uses a hash set and only counts from sequence starts.
// Time O(n) average, space O(n).
func Longest(a []int) int {
	set := make(map[int]bool, len(a))
	for _, x := range a {
		set[x] = true
	}
	best := 0
	for x := range set {
		if set[x-1] {
			continue // not the start of a sequence
		}
		length := 1
		for set[x+length] {
			length++
		}
		best = max(best, length)
	}
	return best
}
