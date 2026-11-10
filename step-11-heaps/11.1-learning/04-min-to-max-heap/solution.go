package mintomaxheap

// ConvertToMaxHeap rearranges a (any order, e.g. a min-heap) into a max-heap
// in place. O(n).
func ConvertToMaxHeap(a []int) {
	n := len(a)
	for i := n/2 - 1; i >= 0; i-- {
		siftDown(a, i, n)
	}
}

func siftDown(a []int, i, n int) {
	for {
		big := i
		l, r := 2*i+1, 2*i+2
		if l < n && a[l] > a[big] {
			big = l
		}
		if r < n && a[r] > a[big] {
			big = r
		}
		if big == i {
			return
		}
		a[i], a[big] = a[big], a[i]
		i = big
	}
}
