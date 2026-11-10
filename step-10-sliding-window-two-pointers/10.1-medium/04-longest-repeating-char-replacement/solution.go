package longestrepeatingcharreplacement

// CharacterReplacement returns the longest uniform substring after at most k
// changes (uppercase A-Z). Time O(n), space O(1).
func CharacterReplacement(s string, k int) int {
	var count [26]int
	best, l, maxFreq := 0, 0, 0
	for r := 0; r < len(s); r++ {
		count[s[r]-'A']++
		maxFreq = max(maxFreq, count[s[r]-'A'])
		if (r-l+1)-maxFreq > k {
			count[s[l]-'A']--
			l++
		}
		best = max(best, r-l+1)
	}
	return best
}

// CharacterReplacementBrute checks every substring. O(26 n^2).
func CharacterReplacementBrute(s string, k int) int {
	best := 0
	for i := 0; i < len(s); i++ {
		var count [26]int
		mf := 0
		for j := i; j < len(s); j++ {
			count[s[j]-'A']++
			mf = max(mf, count[s[j]-'A'])
			if (j-i+1)-mf <= k {
				best = max(best, j-i+1)
			}
		}
	}
	return best
}
