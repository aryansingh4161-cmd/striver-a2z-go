package minimumplatforms

import "sort"

// MinPlatforms returns the maximum number of trains at the station at once.
// Time O(n log n).
func MinPlatforms(arr, dep []int) int {
	a := append([]int{}, arr...)
	d := append([]int{}, dep...)
	sort.Ints(a)
	sort.Ints(d)
	count, best, i, j := 0, 0, 0, 0
	for i < len(a) {
		if a[i] <= d[j] { // arrival first (ties need separate platforms)
			count++
			i++
			best = max(best, count)
		} else {
			count--
			j++
		}
	}
	return best
}
