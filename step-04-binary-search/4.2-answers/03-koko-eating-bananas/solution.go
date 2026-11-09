package kokoeatingbananas

// MinSpeed returns the minimum eating speed to finish all piles in h hours.
// Binary search on the answer. Time O(n log max), space O(1).
func MinSpeed(piles []int, h int) int {
	lo, hi := 1, 0
	for _, p := range piles {
		hi = max(hi, p)
	}
	ans := hi
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if hoursNeeded(piles, mid) <= h {
			ans = mid // works; try slower
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}

func hoursNeeded(piles []int, k int) int {
	total := 0
	for _, p := range piles {
		total += (p + k - 1) / k // ceil(p / k)
	}
	return total
}
