package linearsearch

// Search returns the first index of target in a, or -1. Time O(n).
func Search(a []int, target int) int {
	for i, x := range a {
		if x == target {
			return i
		}
	}
	return -1
}
