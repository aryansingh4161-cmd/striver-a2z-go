package pow

// PowRecursive computes x^n by halving n. Time O(log n), space O(log n).
func PowRecursive(x float64, n int) float64 {
	if n < 0 {
		return 1 / PowRecursive(x, -n)
	}
	if n == 0 {
		return 1
	}
	half := PowRecursive(x, n/2)
	if n%2 == 0 {
		return half * half
	}
	return half * half * x
}

// Pow is the iterative binary exponentiation. Time O(log n), space O(1).
func Pow(x float64, n int) float64 {
	neg := n < 0
	if neg {
		n = -n
	}
	result := 1.0
	for n > 0 {
		if n%2 == 1 {
			result *= x
		}
		x *= x
		n /= 2
	}
	if neg {
		return 1 / result
	}
	return result
}
