package maxpointsfromcards

// MaxScore picks k cards from either end to maximise the sum. Time O(k), space O(1).
func MaxScore(cards []int, k int) int {
	n := len(cards)
	sum := 0
	for i := 0; i < k; i++ {
		sum += cards[i]
	}
	best := sum
	for i := 1; i <= k; i++ {
		sum -= cards[k-i] // give back the last front card
		sum += cards[n-i] // take one more from the back
		best = max(best, sum)
	}
	return best
}
