package reversenumber

import "math"

// Reverse returns n with its digits reversed. Time O(log n), space O(1).
func Reverse(n int) int {
	sign := 1
	if n < 0 {
		sign, n = -1, -n
	}
	rev := 0
	for n > 0 {
		digit := n % 10
		rev = rev*10 + digit
		n /= 10
	}
	return sign * rev
}

// Reverse32 is the LeetCode 7 version: return 0 if the answer
// does not fit in a signed 32-bit integer.
func Reverse32(n int) int {
	r := Reverse(n)
	if r > math.MaxInt32 || r < math.MinInt32 {
		return 0
	}
	return r
}
