package sieve

// Sieve returns isPrime for 0..n. Time O(n log log n), space O(n).
func Sieve(n int) []bool {
	if n < 0 {
		return nil
	}
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for p := 2; p*p <= n; p++ {
		if isPrime[p] {
			for m := p * p; m <= n; m += p {
				isPrime[m] = false
			}
		}
	}
	return isPrime
}

// Primes lists all primes <= n.
func Primes(n int) []int {
	res := []int{}
	for i, ok := range Sieve(n) {
		if ok {
			res = append(res, i)
		}
	}
	return res
}

// CountPrimes counts primes strictly less than n (LeetCode 204).
func CountPrimes(n int) int {
	if n < 3 {
		return 0
	}
	return len(Primes(n - 1))
}
