package mergesort

// Sort sorts a in place using merge sort. Time O(n log n), space O(n).
func Sort(a []int) {
	if len(a) < 2 {
		return
	}
	mergeSort(a, 0, len(a)-1)
}

func mergeSort(a []int, lo, hi int) {
	if lo >= hi {
		return // one element is already sorted
	}
	mid := (lo + hi) / 2
	mergeSort(a, lo, mid)
	mergeSort(a, mid+1, hi)
	merge(a, lo, mid, hi)
}

// merge combines sorted a[lo..mid] and a[mid+1..hi].
func merge(a []int, lo, mid, hi int) {
	temp := make([]int, 0, hi-lo+1)
	left, right := lo, mid+1
	for left <= mid && right <= hi {
		if a[left] <= a[right] { // <= keeps the sort stable
			temp = append(temp, a[left])
			left++
		} else {
			temp = append(temp, a[right])
			right++
		}
	}
	temp = append(temp, a[left:mid+1]...) // leftovers from the left half
	temp = append(temp, a[right:hi+1]...) // leftovers from the right half
	copy(a[lo:hi+1], temp)
}
