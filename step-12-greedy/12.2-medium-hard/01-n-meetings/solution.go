package nmeetings

import "sort"

// MaxMeetings returns the 1-based indexes of the chosen meetings (in the order
// they happen). A meeting must start strictly after the previous one ends.
// Time O(n log n).
func MaxMeetings(start, end []int) []int {
	idx := make([]int, len(start))
	for i := range idx {
		idx[i] = i
	}
	sort.SliceStable(idx, func(a, b int) bool { return end[idx[a]] < end[idx[b]] })
	res := []int{}
	lastEnd := -1 << 62
	for _, i := range idx {
		if start[i] > lastEnd {
			res = append(res, i+1)
			lastEnd = end[i]
		}
	}
	return res
}
