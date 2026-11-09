package printnto1

// CountDown prints n, then recurses with n-1.
func CountDown(n int) []int {
	var out []int
	var f func(k int)
	f = func(k int) {
		if k == 0 {
			return
		}
		out = append(out, k)
		f(k - 1)
	}
	f(n)
	return out
}

// Backtrack goes up to n first, then prints while returning.
func Backtrack(n int) []int {
	var out []int
	var f func(i int)
	f = func(i int) {
		if i > n {
			return
		}
		f(i + 1)
		out = append(out, i)
	}
	f(1)
	return out
}
