package mergeoverlappingintervals

import "sort"

// Merge merges overlapping [start, end] intervals. Time O(n log n).
// The input slice is not modified.
func Merge(intervals [][]int) [][]int {
	iv := make([][]int, len(intervals))
	for i, x := range intervals {
		iv[i] = []int{x[0], x[1]}
	}
	sort.Slice(iv, func(i, j int) bool { return iv[i][0] < iv[j][0] })
	res := [][]int{}
	for _, cur := range iv {
		if len(res) > 0 && cur[0] <= res[len(res)-1][1] {
			last := res[len(res)-1]
			last[1] = max(last[1], cur[1])
		} else {
			res = append(res, cur)
		}
	}
	return res
}
