package atoi

import "math"

// MyAtoi converts s to a clamped 32-bit integer following LeetCode 8 rules.
// Time O(n), space O(1).
func MyAtoi(s string) int {
	i, n := 0, len(s)
	for i < n && s[i] == ' ' {
		i++
	}
	sign := 1
	if i < n && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	result := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		result = result*10 + int(s[i]-'0')
		if sign*result > math.MaxInt32 {
			return math.MaxInt32
		}
		if sign*result < math.MinInt32 {
			return math.MinInt32
		}
		i++
	}
	return sign * result
}
