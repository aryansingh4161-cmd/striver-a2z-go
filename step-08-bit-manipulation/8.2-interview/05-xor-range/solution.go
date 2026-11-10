package xorrange

// XorUpTo returns 1 ^ 2 ^ ... ^ n (0 for n <= 0). O(1).
func XorUpTo(n int) int {
	if n <= 0 {
		return 0
	}
	switch n % 4 {
	case 0:
		return n
	case 1:
		return 1
	case 2:
		return n + 1
	default:
		return 0
	}
}

// XorRange returns L ^ (L+1) ^ ... ^ R for 0 <= L <= R. O(1).
func XorRange(l, r int) int {
	return XorUpTo(r) ^ XorUpTo(l-1)
}
