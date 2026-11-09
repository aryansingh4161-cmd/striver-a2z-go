package mindaysbouquets

// MinDays returns the earliest day to make m bouquets of k adjacent flowers,
// or -1. Time O(n log(max-min)), space O(1).
func MinDays(bloom []int, m, k int) int {
	if m*k > len(bloom) {
		return -1
	}
	lo, hi := bloom[0], bloom[0]
	for _, d := range bloom {
		lo, hi = min(lo, d), max(hi, d)
	}
	ans := -1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if bouquets(bloom, mid, k) >= m {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}

func bouquets(bloom []int, day, k int) int {
	count, made := 0, 0
	for _, d := range bloom {
		if d <= day {
			count++
			if count == k {
				made++
				count = 0
			}
		} else {
			count = 0
		}
	}
	return made
}
