package gasstation

import "math"

// MinMaxDistance returns the smallest possible maximum gap after adding k
// stations. Binary search on a real-valued answer, 100 halving steps.
// Time O(n * 100), space O(1).
func MinMaxDistance(stations []int, k int) float64 {
	lo, hi := 0.0, 0.0
	for i := 1; i < len(stations); i++ {
		hi = math.Max(hi, float64(stations[i]-stations[i-1]))
	}
	for iter := 0; iter < 100; iter++ {
		mid := (lo + hi) / 2
		if mid == 0 {
			lo = mid
			continue
		}
		if stationsNeeded(stations, mid) <= k {
			hi = mid // works; try smaller
		} else {
			lo = mid
		}
	}
	return hi
}

func stationsNeeded(stations []int, d float64) int {
	total := 0
	for i := 1; i < len(stations); i++ {
		gap := float64(stations[i] - stations[i-1])
		total += int(math.Ceil(gap/d)) - 1
	}
	return total
}
