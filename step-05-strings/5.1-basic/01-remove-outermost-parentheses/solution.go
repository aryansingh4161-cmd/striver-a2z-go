package removeoutermostparentheses

import "strings"

// RemoveOuter removes the outermost pair of every primitive part.
// Time O(n), space O(n).
func RemoveOuter(s string) string {
	var sb strings.Builder
	depth := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			if depth > 0 {
				sb.WriteByte('(')
			}
			depth++
		} else {
			depth--
			if depth > 0 {
				sb.WriteByte(')')
			}
		}
	}
	return sb.String()
}
