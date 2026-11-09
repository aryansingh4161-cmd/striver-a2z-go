package countdigits

import "math"

// CountDigits returns the number of digits in n. Time O(log n), space O(1).
func CountDigits(n int) int {
	if n == 0 {
		return 1 // special case: the loop below would return 0
	}
	if n < 0 {
		n = -n
	}
	count := 0
	for n > 0 {
		count++
		n /= 10 // drop the last digit
	}
	return count
}

// CountDigitsLog uses the formula floor(log10(n)) + 1.
func CountDigitsLog(n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
	}
	return int(math.Log10(float64(n))) + 1
}
