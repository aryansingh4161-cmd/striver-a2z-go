package bubblesort

// Sort sorts a in place using bubble sort with the early-exit optimisation.
// Time O(n^2) worst, O(n) if already sorted. Space O(1).
func Sort(a []int) {
	for end := len(a) - 1; end >= 1; end-- {
		swapped := false
		for j := 0; j < end; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}
