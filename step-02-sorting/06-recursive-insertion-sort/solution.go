package recursiveinsertionsort

// Sort sorts a in place with recursive insertion sort.
func Sort(a []int) {
	insert(a, 1)
}

func insert(a []int, i int) {
	if i >= len(a) {
		return
	}
	for j := i; j > 0 && a[j-1] > a[j]; j-- {
		a[j-1], a[j] = a[j], a[j-1]
	}
	insert(a, i+1)
}
