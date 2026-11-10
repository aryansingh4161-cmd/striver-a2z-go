package minimumwindowsubstring

// MinWindow returns the shortest substring of s containing all of t. O(n + m).
func MinWindow(s, t string) string {
	if len(t) == 0 {
		return ""
	}
	var need [256]int
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	missing := len(t)
	bestLen, bestStart := len(s)+1, 0
	l := 0
	for r := 0; r < len(s); r++ {
		if need[s[r]] > 0 {
			missing--
		}
		need[s[r]]--
		for missing == 0 { // valid window: try to shrink
			if r-l+1 < bestLen {
				bestLen, bestStart = r-l+1, l
			}
			need[s[l]]++
			if need[s[l]] > 0 {
				missing++
			}
			l++
		}
	}
	if bestLen > len(s) {
		return ""
	}
	return s[bestStart : bestStart+bestLen]
}
