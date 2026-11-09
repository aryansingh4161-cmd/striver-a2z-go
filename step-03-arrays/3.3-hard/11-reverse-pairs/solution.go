package reversepairs

// CountBrute checks every pair. Time O(n^2).
func CountBrute(a []int) int {
	c := 0
	for i := range a {
		for j := i + 1; j < len(a); j++ {
			if a[i] > 2*a[j] {
				c++
			}
		}
	}
	return c
}

// Count counts reverse pairs with merge sort. Time O(n log n), space O(n).
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
	// counting pass on two sorted halves
	r := mid + 1
	for l := lo; l <= mid; l++ {
		for r <= hi && a[l] > 2*a[r] {
			r++
		}
		c += r - (mid + 1)
	}
	// normal merge
	temp := make([]int, 0, hi-lo+1)
	l, rr := lo, mid+1
	for l <= mid && rr <= hi {
		if a[l] <= a[rr] {
			temp = append(temp, a[l])
			l++
		} else {
			temp = append(temp, a[rr])
			rr++
		}
	}
	temp = append(temp, a[l:mid+1]...)
	temp = append(temp, a[rr:hi+1]...)
	copy(a[lo:hi+1], temp)
	return c
}
