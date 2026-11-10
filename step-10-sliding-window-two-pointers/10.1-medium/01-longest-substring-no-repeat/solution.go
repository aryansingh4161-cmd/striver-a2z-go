package longestsubstringnorepeat

// LongestBrute extends from every start. O(n^2).
func LongestBrute(s string) int {
	best := 0
	for i := 0; i < len(s); i++ {
		var seen [256]bool
		for j := i; j < len(s); j++ {
			if seen[s[j]] {
				break
			}
			seen[s[j]] = true
			best = max(best, j-i+1)
		}
	}
	return best
}

// Longest uses a sliding window with last-seen positions. Time O(n), space O(1).
func Longest(s string) int {
	var last [256]int
	for i := range last {
		last[i] = -1
	}
	best, l := 0, 0
	for r := 0; r < len(s); r++ {
		if last[s[r]] >= l {
			l = last[s[r]] + 1 // jump past the previous copy
		}
		last[s[r]] = r
		best = max(best, r-l+1)
	}
	return best
}
