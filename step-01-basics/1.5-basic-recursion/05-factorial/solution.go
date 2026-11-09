package factorial

// Factorial returns n! (valid for 0 <= n <= 20). Time O(n), space O(n).
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}
