package poweroftwo

// IsPowerOfTwoLoop divides by 2 while even. O(log n).
func IsPowerOfTwoLoop(n int) bool {
	if n <= 0 {
		return false
	}
	for n%2 == 0 {
		n /= 2
	}
	return n == 1
}

// IsPowerOfTwo uses n & (n-1). O(1).
func IsPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
