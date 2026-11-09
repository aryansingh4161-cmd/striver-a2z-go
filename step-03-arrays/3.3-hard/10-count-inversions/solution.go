package countinversions

// CountBrute checks every pair. Time O(n^2).
func CountBrute(a []int) int {
	c := 0
	for i := range a {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				c++
			}
		}
	}
	return c
}

// Count counts inversions with merge sort. Time O(n log n), space O(n).
// It works on a copy, so a is not changed.
func Count(a []int) int {
	b := append([]int{}, a...)
	return sortCount(b, 0, len(b)-1)
}

func sortCount(a []int, lo, hi int) int {
	if lo >= hi {
		return 0
	}
	mid := (lo + hi) / 2
	c := sortCount(a, lo, mid) + sortCount(a, mid+1, hi)
	temp := make([]int, 0, hi-lo+1)
	l, r := lo, mid+1
	for l <= mid && r <= hi {
		if a[l] <= a[r] {
			temp = append(temp, a[l])
			l++
		} else {
			c += mid - l + 1 // a[l..mid] are all bigger than a[r]
			temp = append(temp, a[r])
			r++
		}
	}
	temp = append(temp, a[l:mid+1]...)
	temp = append(temp, a[r:hi+1]...)
	copy(a[lo:hi+1], temp)
	return c
}
