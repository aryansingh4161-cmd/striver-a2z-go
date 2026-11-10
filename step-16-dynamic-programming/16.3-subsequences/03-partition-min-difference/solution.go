package partitionmindifference

// MinDifference returns the smallest |sum1 − sum2| over all splits.
// Uses the subset-sum row for every target. O(n·total).
func MinDifference(arr []int) int {
	total := 0
	for _, x := range arr {
		total += x
	}
	can := make([]bool, total+1) // can[s]: some subset sums to s
	can[0] = true
	for _, x := range arr {
		for s := total; s >= x; s-- {
			can[s] = can[s] || can[s-x]
		}
	}
	best := total
	for s1 := 0; s1 <= total/2; s1++ {
		if can[s1] {
			best = min(best, total-2*s1)
		}
	}
	return best
}
