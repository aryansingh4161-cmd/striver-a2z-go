package sortcharactersbyfrequency

import (
	"sort"
	"strings"
)

// FrequencySort orders characters by decreasing frequency (ties: smaller byte first).
// Time O(n + k log k), space O(n).
func FrequencySort(s string) string {
	var count [256]int
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}
	var chars []byte
	for c := 0; c < 256; c++ {
		if count[c] > 0 {
			chars = append(chars, byte(c))
		}
	}
	sort.Slice(chars, func(i, j int) bool {
		if count[chars[i]] != count[chars[j]] {
			return count[chars[i]] > count[chars[j]]
		}
		return chars[i] < chars[j]
	})
	var sb strings.Builder
	for _, c := range chars {
		sb.WriteString(strings.Repeat(string(c), count[c]))
	}
	return sb.String()
}
