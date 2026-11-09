package gcd

// GCDBrute tries every candidate from min(a,b) down. Time O(min(a,b)).
func GCDBrute(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	for i := min(a, b); i >= 1; i-- {
		if a%i == 0 && b%i == 0 {
			return i
		}
	}
	return 1
}

// GCD uses the Euclidean algorithm. Time O(log(min(a,b))), space O(1).
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM returns the least common multiple of a and b.
func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return a / GCD(a, b) * b
}
