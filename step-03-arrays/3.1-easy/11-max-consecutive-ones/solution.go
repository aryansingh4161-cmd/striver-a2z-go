package maxconsecutiveones

// MaxOnes returns the length of the longest run of 1s. Time O(n), space O(1).
func MaxOnes(a []int) int {
	best, count := 0, 0
	for _, x := range a {
		if x == 1 {
			count++
			best = max(best, count)
		} else {
			count = 0
		}
	}
	return best
}
