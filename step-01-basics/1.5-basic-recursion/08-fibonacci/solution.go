package fibonacci

// FibRecursive is the plain two-call recursion. Time O(2^n). Only for small n.
func FibRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return FibRecursive(n-1) + FibRecursive(n-2)
}

// Fib keeps only the last two values. Time O(n), space O(1).
func Fib(n int) int {
	if n <= 1 {
		return n
	}
	prev2, prev1 := 0, 1
	for i := 2; i <= n; i++ {
		prev2, prev1 = prev1, prev1+prev2
	}
	return prev1
}
