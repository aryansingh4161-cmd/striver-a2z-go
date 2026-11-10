package recursiveatoi

import "math"

// MyAtoi parses s like LeetCode 8, reading digits recursively.
// Time O(n), space O(n) stack.
func MyAtoi(s string) int {
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	sign := 1
	if i < len(s) && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	var digits func(i, acc int) int
	digits = func(i, acc int) int {
		if i >= len(s) || s[i] < '0' || s[i] > '9' {
			return acc
		}
		acc = acc*10 + int(s[i]-'0')
		if acc > math.MaxInt32+1 { // already too big; stop early
			return acc
		}
		return digits(i+1, acc)
	}
	v := sign * digits(i, 0)
	if v > math.MaxInt32 {
		return math.MaxInt32
	}
	if v < math.MinInt32 {
		return math.MinInt32
	}
	return v
}
