package countsetbits

// CountShift checks the last bit and shifts. O(log n). n must be >= 0.
func CountShift(n int) int {
	c := 0
	for n > 0 {
		c += n & 1
		n >>= 1
	}
	return c
}

// CountKernighan removes the lowest set bit each step. O(set bits).
func CountKernighan(n int) int {
	c := 0
	for n != 0 {
		n &= n - 1
		c++
	}
	return c
}

// CountAll returns set-bit counts for 0..n. O(n).
func CountAll(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		res[i] = res[i>>1] + i&1
	}
	return res
}
