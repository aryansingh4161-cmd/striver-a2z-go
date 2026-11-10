package oddeven

// IsOdd uses the last bit. Works for negative numbers. O(1).
func IsOdd(n int) bool { return n&1 == 1 }
