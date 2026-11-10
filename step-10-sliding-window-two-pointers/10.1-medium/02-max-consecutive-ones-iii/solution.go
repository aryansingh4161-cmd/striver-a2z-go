package maxconsecutiveonesiii

// LongestOnes uses a shrinking window: at most k zeros inside. Time O(n).
func LongestOnes(a []int, k int) int {
	best, l, zeros := 0, 0, 0
	for r := range a {
		if a[r] == 0 {
			zeros++
		}
		for zeros > k {
			if a[l] == 0 {
				zeros--
			}
			l++
		}
		best = max(best, r-l+1)
	}
	return best
}

// LongestOnesNoShrink never shrinks the window, only slides it. Time O(n).
func LongestOnesNoShrink(a []int, k int) int {
	l, zeros := 0, 0
	for r := range a {
		if a[r] == 0 {
			zeros++
		}
		if zeros > k {
			if a[l] == 0 {
				zeros--
			}
			l++
		}
	}
	return len(a) - l
}
