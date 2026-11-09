package checkprime

// IsPrimeBrute counts all divisors. Time O(n).
func IsPrimeBrute(n int) bool {
	if n < 2 {
		return false
	}
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count == 2
}

// IsPrime checks divisors only up to sqrt(n). Time O(sqrt(n)), space O(1).
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
