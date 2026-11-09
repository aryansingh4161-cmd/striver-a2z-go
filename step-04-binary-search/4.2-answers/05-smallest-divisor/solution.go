package smallestdivisor

// SmallestDivisor returns the smallest d whose sum of ceil(a[i]/d) is <= threshold.
// Time O(n log max), space O(1).
func SmallestDivisor(a []int, threshold int) int {
	lo, hi := 1, 0
	for _, x := range a {
		hi = max(hi, x)
	}
	ans := hi
	for lo <= hi {
		mid := lo + (hi-lo)/2
		sum := 0
		for _, x := range a {
			sum += (x + mid - 1) / mid
		}
		if sum <= threshold {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}
