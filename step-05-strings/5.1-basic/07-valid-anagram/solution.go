package validanagram

import (
	"slices"
)

// IsAnagramSort sorts the bytes of both strings. Time O(n log n).
func IsAnagramSort(s, t string) bool {
	a, b := []byte(s), []byte(t)
	slices.Sort(a)
	slices.Sort(b)
	return string(a) == string(b)
}

// IsAnagram counts letters. Time O(n), space O(1). Lowercase a-z only.
func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var count [26]int
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}
	for _, c := range count {
		if c != 0 {
			return false
		}
	}
	return true
}
