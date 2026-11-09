package quicksort

// Sort sorts a in place using quick sort (first element as pivot).
// Average time O(n log n), worst O(n^2). Space O(log n) stack on average.
func Sort(a []int) {
	quickSort(a, 0, len(a)-1)
}

func quickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

// partition places a[lo] (the pivot) at its sorted position and returns that index.
func partition(a []int, lo, hi int) int {
	pivot := a[lo]
	i, j := lo, hi
	for i < j {
		for a[i] <= pivot && i <= hi-1 {
			i++
		}
		for a[j] > pivot && j >= lo+1 {
			j--
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}
	a[lo], a[j] = a[j], a[lo]
	return j
}
