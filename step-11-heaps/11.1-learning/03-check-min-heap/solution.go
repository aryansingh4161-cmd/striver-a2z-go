package checkminheap

// IsMinHeap reports whether a satisfies the min-heap property. O(n).
func IsMinHeap(a []int) bool {
	n := len(a)
	for i := 0; i < n/2; i++ {
		l, r := 2*i+1, 2*i+2
		if a[i] > a[l] || (r < n && a[i] > a[r]) {
			return false
		}
	}
	return true
}
