package largestelement

import "sort"

// LargestSort sorts a copy and returns the last element. Time O(n log n).
func LargestSort(a []int) int {
	b := append([]int{}, a...)
	sort.Ints(b)
	return b[len(b)-1]
}

// Largest scans once, keeping the biggest seen so far. Time O(n), space O(1).
// a must be non-empty.
func Largest(a []int) int {
	largest := a[0]
	for _, x := range a {
		if x > largest {
			largest = x
		}
	}
	return largest
}
