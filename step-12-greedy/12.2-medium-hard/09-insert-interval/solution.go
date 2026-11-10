package insertinterval

// Insert adds iv to sorted non-overlapping intervals, merging overlaps. O(n).
func Insert(intervals [][]int, iv []int) [][]int {
	res := [][]int{}
	cur := []int{iv[0], iv[1]}
	i, n := 0, len(intervals)
	for i < n && intervals[i][1] < cur[0] { // entirely before
		res = append(res, intervals[i])
		i++
	}
	for i < n && intervals[i][0] <= cur[1] { // overlapping
		cur[0] = min(cur[0], intervals[i][0])
		cur[1] = max(cur[1], intervals[i][1])
		i++
	}
	res = append(res, cur)
	res = append(res, intervals[i:]...) // entirely after
	return res
}
