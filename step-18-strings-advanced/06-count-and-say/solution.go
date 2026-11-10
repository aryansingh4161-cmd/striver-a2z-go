package countandsay

import (
	"strconv"
	"strings"
)

// next reads s aloud: run-length encode it as count + digit.
func next(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); {
		j := i
		for j < len(s) && s[j] == s[i] {
			j++
		}
		b.WriteString(strconv.Itoa(j - i))
		b.WriteByte(s[i])
		i = j
	}
	return b.String()
}

// CountAndSay returns the n-th term (n ≥ 1).
func CountAndSay(n int) string {
	cur := "1"
	for k := 1; k < n; k++ {
		cur = next(cur)
	}
	return cur
}
