package print1ton

// CountUp prints i, then recurses with i+1.
func CountUp(n int) []int {
	var out []int
	var f func(i int)
	f = func(i int) {
		if i > n {
			return
		}
		out = append(out, i)
		f(i + 1)
	}
	f(1)
	return out
}

// Backtrack recurses first, then prints on the way back.
func Backtrack(n int) []int {
	var out []int
	var f func(k int)
	f = func(k int) {
		if k == 0 {
			return
		}
		f(k - 1)
		out = append(out, k) // runs after the smaller call returns
	}
	f(n)
	return out
}
