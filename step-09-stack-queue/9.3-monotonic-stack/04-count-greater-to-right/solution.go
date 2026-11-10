package countgreatertoright

// CountQueries scans to the right for each query. O(n * q).
func CountQueries(a []int, queries []int) []int {
	res := make([]int, len(queries))
	for k, i := range queries {
		for j := i + 1; j < len(a); j++ {
			if a[j] > a[i] {
				res[k]++
			}
		}
	}
	return res
}

// CountAll returns, for every index, how many elements to its right are greater.
// Merge sort on indexes. Time O(n log n), space O(n).
func CountAll(a []int) []int {
	n := len(a)
	cnt := make([]int, n)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	var sortRange func(lo, hi int)
	sortRange = func(lo, hi int) {
		if lo >= hi {
			return
		}
		mid := (lo + hi) / 2
		sortRange(lo, mid)
		sortRange(mid+1, hi)
		// both halves sorted by value ascending
		temp := make([]int, 0, hi-lo+1)
		l, r := lo, mid+1
		for l <= mid && r <= hi {
			if a[idx[l]] < a[idx[r]] {
				cnt[idx[l]] += hi - r + 1 // everything from r..hi is bigger
				temp = append(temp, idx[l])
				l++
			} else {
				temp = append(temp, idx[r])
				r++
			}
		}
		temp = append(temp, idx[l:mid+1]...)
		temp = append(temp, idx[r:hi+1]...)
		copy(idx[lo:hi+1], temp)
	}
	sortRange(0, n-1)
	return cnt
}
