package checksorted

// IsSorted reports whether a is in non-decreasing order. Time O(n).
func IsSorted(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			return false
		}
	}
	return true
}

// IsSortedRotated reports whether a is a sorted array rotated some number
// of times (LeetCode 1752). It counts drops, including the wrap-around.
func IsSortedRotated(a []int) bool {
	n := len(a)
	drops := 0
	for i := 0; i < n; i++ {
		if a[i] > a[(i+1)%n] {
			drops++
		}
	}
	return drops <= 1
}
