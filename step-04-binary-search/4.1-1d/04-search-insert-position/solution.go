package searchinsertposition

// SearchInsert returns the index of target, or where it would be inserted.
// It is the lower bound. Time O(log n).
func SearchInsert(a []int, target int) int {
	lo, hi, ans := 0, len(a)-1, len(a)
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if a[mid] >= target {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}
