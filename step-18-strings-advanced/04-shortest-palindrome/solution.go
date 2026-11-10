package shortestpalindrome

import kmplps "github.com/aryansingh4161-cmd/striver-a2z-go/step-18-strings-advanced/03-kmp-lps"

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func isPal(s string) bool { return s == reverse(s) }

// Brute finds the longest palindromic prefix by checking every length. O(n²).
func Brute(s string) string {
	for l := len(s); l >= 0; l-- {
		if isPal(s[:l]) {
			return reverse(s[l:]) + s
		}
	}
	return s
}

// Shortest uses LPS of s + "#" + reverse(s). O(n). s must not contain '#'.
func Shortest(s string) string {
	if s == "" {
		return s
	}
	lps := kmplps.LPS(s + "#" + reverse(s))
	l := lps[len(lps)-1] // longest palindromic prefix
	return reverse(s[l:]) + s
}

// MinCharsToAdd is Striver's version: how many characters must be added in front.
func MinCharsToAdd(s string) int { return len(Shortest(s)) - len(s) }
