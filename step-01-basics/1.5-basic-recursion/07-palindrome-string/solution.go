package palindromestring

import (
	"strings"
	"unicode"
)

// IsPalindrome checks s exactly as given, using recursion.
func IsPalindrome(s string) bool {
	n := len(s)
	var f func(i int) bool
	f = func(i int) bool {
		if i >= n/2 {
			return true
		}
		if s[i] != s[n-1-i] {
			return false
		}
		return f(i + 1)
	}
	return f(0)
}

// IsPalindromeClean is the LeetCode 125 version: only letters and digits count,
// and case is ignored.
func IsPalindromeClean(s string) bool {
	var sb strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			sb.WriteRune(unicode.ToLower(r))
		}
	}
	return IsPalindrome(sb.String())
}
