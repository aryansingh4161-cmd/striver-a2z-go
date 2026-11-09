package longestpalindromicsubstring

// LongestPalindrome expands around every centre. Time O(n^2), space O(1).
func LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	start, best := 0, 1
	expand := func(l, r int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		// the palindrome is s[l+1 .. r-1]
		if r-l-1 > best {
			start, best = l+1, r-l-1
		}
	}
	for i := 0; i < len(s); i++ {
		expand(i, i)   // odd length
		expand(i, i+1) // even length
	}
	return s[start : start+best]
}
