package alldivisors

import "sort"

// DivisorsBrute checks every number from 1 to n. Time O(n).
func DivisorsBrute(n int) []int {
	var res []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			res = append(res, i)
		}
	}
	return res
}

// Divisors uses divisor pairs (i, n/i) with i up to sqrt(n).
// Time O(sqrt(n) + d log d), space O(d).
func Divisors(n int) []int {
	var res []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			res = append(res, i)
			if n/i != i { // avoid adding the square root twice
				res = append(res, n/i)
			}
		}
	}
	sort.Ints(res)
	return res
}
