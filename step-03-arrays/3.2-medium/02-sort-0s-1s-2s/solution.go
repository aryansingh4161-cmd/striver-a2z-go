package sort012

// SortCount counts 0s, 1s and 2s and rewrites the slice. Time O(n), two passes.
func SortCount(a []int) {
	var c [3]int
	for _, x := range a {
		c[x]++
	}
	i := 0
	for v := 0; v < 3; v++ {
		for k := 0; k < c[v]; k++ {
			a[i] = v
			i++
		}
	}
}

// SortDNF uses the Dutch National Flag algorithm. Time O(n), one pass, space O(1).
func SortDNF(a []int) {
	low, mid, high := 0, 0, len(a)-1
	for mid <= high {
		switch a[mid] {
		case 0:
			a[low], a[mid] = a[mid], a[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			a[mid], a[high] = a[high], a[mid]
			high-- // don't move mid: the new a[mid] is unchecked
		}
	}
}
