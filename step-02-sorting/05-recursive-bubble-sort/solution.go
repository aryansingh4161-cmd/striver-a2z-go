package recursivebubblesort

// Sort sorts a in place with recursive bubble sort.
func Sort(a []int) {
	bubble(a, len(a))
}

func bubble(a []int, n int) {
	if n <= 1 {
		return
	}
	swapped := false
	for j := 0; j < n-1; j++ {
		if a[j] > a[j+1] {
			a[j], a[j+1] = a[j+1], a[j]
			swapped = true
		}
	}
	if !swapped {
		return
	}
	bubble(a, n-1) // the largest is now at index n-1
}
