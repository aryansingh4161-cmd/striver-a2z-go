package stockfee

// Recursive: Stock II, paying the fee on every sell. O(2^n).
func Recursive(p []int, fee int) int {
	var f func(i int, canBuy bool) int
	f = func(i int, canBuy bool) int {
		if i == len(p) {
			return 0
		}
		if canBuy {
			return max(-p[i]+f(i+1, false), f(i+1, true))
		}
		return max(p[i]-fee+f(i+1, true), f(i+1, false))
	}
	return f(0, true)
}

// Optimal keeps tomorrow's two values. O(n) time, O(1) space.
func Optimal(p []int, fee int) int {
	aheadBuy, aheadHold := 0, 0
	for i := len(p) - 1; i >= 0; i-- {
		buy := max(-p[i]+aheadHold, aheadBuy)
		hold := max(p[i]-fee+aheadBuy, aheadHold)
		aheadBuy, aheadHold = buy, hold
	}
	return aheadBuy
}
