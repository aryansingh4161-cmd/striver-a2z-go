package shippackages

// ShipWithinDays returns the minimum ship capacity to deliver all packages
// (in order) within days. Time O(n log sum), space O(1).
func ShipWithinDays(w []int, days int) int {
	lo, hi := 0, 0
	for _, x := range w {
		lo = max(lo, x)
		hi += x
	}
	ans := hi
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if daysNeeded(w, mid) <= days {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}

func daysNeeded(w []int, capacity int) int {
	days, load := 1, 0
	for _, x := range w {
		if load+x > capacity {
			days++
			load = 0
		}
		load += x
	}
	return days
}
