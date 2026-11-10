package spffactorisation

// SPF holds smallest-prime-factor values for 0..N.
type SPF []int

// NewSPF precomputes smallest prime factors up to n. Time O(n log log n).
func NewSPF(n int) SPF {
	spf := make(SPF, n+1)
	for i := range spf {
		spf[i] = i
	}
	for p := 2; p*p <= n; p++ {
		if spf[p] == p {
			for m := p * p; m <= n; m += p {
				if spf[m] == m {
					spf[m] = p
				}
			}
		}
	}
	return spf
}

// Factorise returns the prime factors of x with repeats, in increasing order.
// x must be between 1 and N. Time O(log x).
func (spf SPF) Factorise(x int) []int {
	res := []int{}
	for x > 1 {
		res = append(res, spf[x])
		x /= spf[x]
	}
	return res
}
