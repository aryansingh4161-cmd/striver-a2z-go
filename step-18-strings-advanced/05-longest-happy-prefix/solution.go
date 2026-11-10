package longesthappyprefix

import kmplps "github.com/aryansingh4161-cmd/striver-a2z-go/step-18-strings-advanced/03-kmp-lps"

// Brute tries every length from long to short. O(n²).
func Brute(s string) string {
	for l := len(s) - 1; l >= 1; l-- {
		if s[:l] == s[len(s)-l:] {
			return s[:l]
		}
	}
	return ""
}

// LongestPrefix is the last value of the LPS array. O(n).
func LongestPrefix(s string) string {
	if s == "" {
		return ""
	}
	lps := kmplps.LPS(s)
	return s[:lps[len(s)-1]]
}
