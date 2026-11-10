package alldivisors

import "sort"

// Divisors returns all divisors of n in increasing order. Time O(sqrt n).
func Divisors(n int) []int {
	res := []int{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			res = append(res, i)
			if i != n/i {
				res = append(res, n/i)
			}
		}
	}
	sort.Ints(res)
	return res
}
