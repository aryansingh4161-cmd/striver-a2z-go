package bookallocation

// MinMaxPages returns the minimum possible maximum pages per student,
// or -1 if there are fewer books than students. Time O(n log sum).
func MinMaxPages(pages []int, m int) int {
	if m > len(pages) {
		return -1
	}
	return minLargestSplit(pages, m)
}

// minLargestSplit splits a into at most k contiguous parts and minimises
// the largest part sum. Used by book allocation, split array and painters.
func minLargestSplit(a []int, k int) int {
	lo, hi := 0, 0
	for _, x := range a {
		lo = max(lo, x)
		hi += x
	}
	ans := hi
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if partsNeeded(a, mid) <= k {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}

func partsNeeded(a []int, limit int) int {
	parts, sum := 1, 0
	for _, x := range a {
		if sum+x > limit {
			parts++
			sum = 0
		}
		sum += x
	}
	return parts
}
