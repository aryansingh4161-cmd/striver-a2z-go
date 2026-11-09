package sumfirstn

// SumParam passes the running sum as a parameter.
func SumParam(n int) int {
	answer := 0
	var f func(i, sum int)
	f = func(i, sum int) {
		if i < 1 {
			answer = sum
			return
		}
		f(i-1, sum+i)
	}
	f(n, 0)
	return answer
}

// SumFunc returns the answer from each call: n + sum(n-1).
func SumFunc(n int) int {
	if n == 0 {
		return 0
	}
	return n + SumFunc(n-1)
}

// SumFormula uses n(n+1)/2. O(1).
func SumFormula(n int) int {
	return n * (n + 1) / 2
}
