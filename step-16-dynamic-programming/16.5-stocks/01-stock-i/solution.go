package stocki

// Brute tries every buy/sell pair. O(n²).
func Brute(prices []int) int {
	best := 0
	for i := range prices {
		for j := i + 1; j < len(prices); j++ {
			best = max(best, prices[j]-prices[i])
		}
	}
	return best
}

// Optimal keeps the cheapest price seen so far. O(n) time, O(1) space.
func Optimal(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minSoFar, best := prices[0], 0
	for _, p := range prices[1:] {
		best = max(best, p-minSoFar)
		minSoFar = min(minSoFar, p)
	}
	return best
}
