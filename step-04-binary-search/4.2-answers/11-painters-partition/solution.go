package painterspartition

// MinTime returns the minimum time for k painters to paint all boards
// (each painter takes a contiguous block). Time O(n log sum).
func MinTime(boards []int, k int) int {
	lo, hi := 0, 0
	for _, x := range boards {
		lo = max(lo, x)
		hi += x
	}
	ans := hi
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if paintersNeeded(boards, mid) <= k {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}

func paintersNeeded(boards []int, limit int) int {
	painters, load := 1, 0
	for _, b := range boards {
		if load+b > limit {
			painters++
			load = 0
		}
		load += b
	}
	return painters
}
