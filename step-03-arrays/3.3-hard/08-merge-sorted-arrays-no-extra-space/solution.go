package mergesortedarrays

import "sort"

// MergeSwapSort: swap a's largest with b's smallest while out of order, then sort both.
// Space O(1).
func MergeSwapSort(a, b []int) {
	i, j := len(a)-1, 0
	for i >= 0 && j < len(b) && a[i] > b[j] {
		a[i], b[j] = b[j], a[i]
		i--
		j++
	}
	sort.Ints(a)
	sort.Ints(b)
}

// MergeGap uses the gap method over the virtual array a+b. Space O(1).
func MergeGap(a, b []int) {
	n, m := len(a), len(b)
	total := n + m
	if total < 2 {
		return
	}
	// get returns a pointer to element k of the combined array.
	get := func(k int) *int {
		if k < n {
			return &a[k]
		}
		return &b[k-n]
	}
	gap := (total + 1) / 2
	for {
		for l := 0; l+gap < total; l++ {
			x, y := get(l), get(l+gap)
			if *x > *y {
				*x, *y = *y, *x
			}
		}
		if gap == 1 {
			break
		}
		gap = (gap + 1) / 2
	}
}

// MergeIntoFirst is LeetCode 88: a has len m+n with the last n slots empty;
// the first m are real values. Fills from the back. Time O(m+n), space O(1).
func MergeIntoFirst(a []int, m int, b []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for j >= 0 {
		if i >= 0 && a[i] > b[j] {
			a[k] = a[i]
			i--
		} else {
			a[k] = b[j]
			j--
		}
		k--
	}
}
