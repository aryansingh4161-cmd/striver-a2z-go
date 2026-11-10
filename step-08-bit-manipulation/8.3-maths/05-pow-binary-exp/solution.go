package powbinaryexp

// Pow computes x^n for n >= 0 (any sign of n is handled). O(log n).
func Pow(x float64, n int) float64 {
	neg := n < 0
	if neg {
		n = -n
	}
	result := 1.0
	for n > 0 {
		if n&1 == 1 {
			result *= x
		}
		x *= x
		n >>= 1
	}
	if neg {
		return 1 / result
	}
	return result
}

// PowMod computes base^exp % mod for exp >= 0. O(log exp).
func PowMod(base, exp, mod int) int {
	base %= mod
	result := 1 % mod
	for exp > 0 {
		if exp&1 == 1 {
			result = result * base % mod
		}
		base = base * base % mod
		exp >>= 1
	}
	return result
}
