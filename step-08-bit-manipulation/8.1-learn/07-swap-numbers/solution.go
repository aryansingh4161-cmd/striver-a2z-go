package swapnumbers

// SwapXOR swaps two ints using XOR (for learning; in Go just use a, b = b, a).
func SwapXOR(a, b int) (int, int) {
	a ^= b
	b ^= a // b = original a
	a ^= b // a = original b
	return a, b
}
