package stockbuysell

// MaxProfit returns the best profit from one buy followed by one sell.
// Time O(n), space O(1).
func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice, best := prices[0], 0
	for _, p := range prices {
		best = max(best, p-minPrice)
		minPrice = min(minPrice, p)
	}
	return best
}
