package stockcooldown

// Recursive: selling on day i jumps to day i+2. Exponential.
func Recursive(p []int) int {
	var f func(i int, canBuy bool) int
	f = func(i int, canBuy bool) int {
		if i >= len(p) {
			return 0
		}
		if canBuy {
			return max(-p[i]+f(i+1, false), f(i+1, true))
		}
		return max(p[i]+f(i+2, true), f(i+1, false))
	}
	return f(0, true)
}

// Tab uses two extra zero rows so i+2 is always in range. O(n).
func Tab(p []int) int {
	n := len(p)
	buy, hold := make([]int, n+2), make([]int, n+2)
	for i := n - 1; i >= 0; i-- {
		buy[i] = max(-p[i]+hold[i+1], buy[i+1])
		hold[i] = max(p[i]+buy[i+2], hold[i+1])
	}
	return buy[0]
}

// Optimal keeps days i+1 and i+2. O(1) space.
func Optimal(p []int) int {
	buy1, hold1, buy2 := 0, 0, 0 // buy1/hold1 = day i+1, buy2 = day i+2
	for i := len(p) - 1; i >= 0; i-- {
		buy := max(-p[i]+hold1, buy1)
		hold := max(p[i]+buy2, hold1)
		buy2 = buy1
		buy1, hold1 = buy, hold
	}
	return buy1
}
