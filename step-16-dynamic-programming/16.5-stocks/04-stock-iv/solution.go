package stockiv

// Recursive: state (day, canBuy, cap). O(2^n).
func Recursive(k int, p []int) int {
	var f func(i int, canBuy bool, cap int) int
	f = func(i int, canBuy bool, cap int) int {
		if i == len(p) || cap == 0 {
			return 0
		}
		if canBuy {
			return max(-p[i]+f(i+1, false, cap), f(i+1, true, cap))
		}
		return max(p[i]+f(i+1, true, cap-1), f(i+1, false, cap))
	}
	return f(0, true, k)
}

// Optimal: Stock III with k instead of 2. O(n·k) time, O(k) space.
func Optimal(k int, p []int) int {
	aheadBuy, aheadHold := make([]int, k+1), make([]int, k+1)
	for i := len(p) - 1; i >= 0; i-- {
		curBuy, curHold := make([]int, k+1), make([]int, k+1)
		for cap := 1; cap <= k; cap++ {
			curBuy[cap] = max(-p[i]+aheadHold[cap], aheadBuy[cap])
			curHold[cap] = max(p[i]+aheadBuy[cap-1], aheadHold[cap])
		}
		aheadBuy, aheadHold = curBuy, curHold
	}
	return aheadBuy[k]
}

// TransactionNumber uses f(i, t) with t = 0..2k (even = buy, odd = sell). O(n·k).
func TransactionNumber(k int, p []int) int {
	ahead := make([]int, 2*k+1)
	for i := len(p) - 1; i >= 0; i-- {
		cur := make([]int, 2*k+1)
		for t := 2*k - 1; t >= 0; t-- {
			if t%2 == 0 {
				cur[t] = max(-p[i]+ahead[t+1], ahead[t])
			} else {
				cur[t] = max(p[i]+ahead[t+1], ahead[t])
			}
		}
		ahead = cur
	}
	return ahead[0]
}
