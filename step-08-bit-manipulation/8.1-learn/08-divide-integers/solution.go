package divideintegers

import "math"

// Divide returns dividend / divisor truncated toward zero, using only
// shifts, addition and subtraction. Clamped to 32-bit range. Time O(log^2 n).
func Divide(dividend, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	neg := (dividend < 0) != (divisor < 0)
	a, b := abs(dividend), abs(divisor)
	q := 0
	for a >= b {
		k := 0
		for a >= b<<(k+1) {
			k++
		}
		a -= b << k
		q += 1 << k
	}
	if neg {
		q = -q
	}
	return q
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
