package primefactors

// PrimeFactors returns the distinct prime factors of n in increasing order.
// Time O(sqrt(n)).
func PrimeFactors(n int) []int {
	res := []int{}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			res = append(res, i)
			for n%i == 0 {
				n /= i
			}
		}
	}
	if n > 1 {
		res = append(res, n)
	}
	return res
}
