package mergeintervals

import "sort"

// Merge merges overlapping intervals. Time O(n log n).
func Merge(intervals [][]int) [][]int {
	iv := make([][]int, len(intervals))
	for i, x := range intervals {
		iv[i] = []int{x[0], x[1]}
	}
	sort.Slice(iv, func(i, j int) bool { return iv[i][0] < iv[j][0] })
	res := [][]int{}
	for _, cur := range iv {
		if n := len(res); n > 0 && cur[0] <= res[n-1][1] {
			res[n-1][1] = max(res[n-1][1], cur[1])
		} else {
			res = append(res, cur)
		}
	}
	return res
}
