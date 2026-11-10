package stockiii

// Recursive: state (day, canBuy, transactions left). O(2^n).
func Recursive(p []int) int {
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
	return f(0, true, 2)
}

// Optimal: ahead[canBuy][cap] for the next day only. O(n) time, O(1) space.
func Optimal(p []int) int {
	var ahead [2][3]int // [0] = holding, [1] = can buy
	for i := len(p) - 1; i >= 0; i-- {
		var cur [2][3]int
		for cap := 1; cap <= 2; cap++ {
			cur[1][cap] = max(-p[i]+ahead[0][cap], ahead[1][cap])
			cur[0][cap] = max(p[i]+ahead[1][cap-1], ahead[0][cap])
		}
		ahead = cur
	}
	return ahead[1][2]
}

// FourStates is the forward version with buy1, sell1, buy2, sell2.
func FourStates(p []int) int {
	const negInf = -(1 << 60)
	buy1, sell1, buy2, sell2 := negInf, 0, negInf, 0
	for _, x := range p {
		buy1 = max(buy1, -x)
		sell1 = max(sell1, buy1+x)
		buy2 = max(buy2, sell1-x)
		sell2 = max(sell2, buy2+x)
	}
	return sell2
}
