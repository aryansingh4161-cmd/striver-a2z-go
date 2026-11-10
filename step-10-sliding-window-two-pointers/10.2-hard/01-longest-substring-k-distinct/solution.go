package longestsubstringkdistinct

// Longest returns the longest substring with at most k distinct characters. O(n).
func Longest(s string, k int) int {
	count := map[byte]int{}
	best, l := 0, 0
	for r := 0; r < len(s); r++ {
		count[s[r]]++
		for len(count) > k {
			count[s[l]]--
			if count[s[l]] == 0 {
				delete(count, s[l])
			}
			l++
		}
		best = max(best, r-l+1)
	}
	return best
}

// LongestBrute extends from every start. O(n^2).
func LongestBrute(s string, k int) int {
	best := 0
	for i := 0; i < len(s); i++ {
		seen := map[byte]bool{}
		for j := i; j < len(s); j++ {
			seen[s[j]] = true
			if len(seen) > k {
				break
			}
			best = max(best, j-i+1)
		}
	}
	return best
}
