package selectionsort

// Sort sorts a in place using selection sort. Time O(n^2), space O(1).
func Sort(a []int) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[minIdx] {
				minIdx = j
			}
		}
		a[i], a[minIdx] = a[minIdx], a[i]
	}
}
