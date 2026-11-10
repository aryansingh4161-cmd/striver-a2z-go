package nonoverlappingintervals

import "sort"

// EraseOverlap returns the minimum intervals to remove so none overlap
// (touching ends allowed). Time O(n log n).
func EraseOverlap(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	iv := append([][]int{}, intervals...)
	sort.Slice(iv, func(i, j int) bool { return iv[i][1] < iv[j][1] })
	kept, lastEnd := 1, iv[0][1]
	for _, x := range iv[1:] {
		if x[0] >= lastEnd {
			kept++
			lastEnd = x[1]
		}
	}
	return len(iv) - kept
}
