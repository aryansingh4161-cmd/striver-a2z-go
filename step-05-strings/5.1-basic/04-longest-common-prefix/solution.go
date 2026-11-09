package longestcommonprefix

import "sort"

// LCP scans column by column. Time O(total chars), space O(1).
func LCP(words []string) string {
	if len(words) == 0 {
		return ""
	}
	first := words[0]
	for i := 0; i < len(first); i++ {
		for _, w := range words[1:] {
			if i >= len(w) || w[i] != first[i] {
				return first[:i]
			}
		}
	}
	return first
}

// LCPSort sorts a copy and compares the first and last word.
func LCPSort(words []string) string {
	if len(words) == 0 {
		return ""
	}
	w := append([]string{}, words...)
	sort.Strings(w)
	a, b := w[0], w[len(w)-1]
	i := 0
	for i < len(a) && i < len(b) && a[i] == b[i] {
		i++
	}
	return a[:i]
}
