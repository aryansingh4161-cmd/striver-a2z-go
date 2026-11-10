package substringsallthreechars

// NumberOfSubstrings counts substrings containing a, b and c. Time O(n), space O(1).
func NumberOfSubstrings(s string) int {
	last := [3]int{-1, -1, -1}
	count := 0
	for r := 0; r < len(s); r++ {
		last[s[r]-'a'] = r
		count += min(last[0], last[1], last[2]) + 1
	}
	return count
}

// NumberOfSubstringsBrute extends from each start until all three appear. O(n^2).
func NumberOfSubstringsBrute(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		var seen [3]bool
		for j := i; j < len(s); j++ {
			seen[s[j]-'a'] = true
			if seen[0] && seen[1] && seen[2] {
				count += len(s) - j
				break
			}
		}
	}
	return count
}
