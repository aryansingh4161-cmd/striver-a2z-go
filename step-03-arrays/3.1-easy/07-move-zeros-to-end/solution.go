package movezerostoend

// MoveZeros moves all zeros to the end, keeping the order of other elements.
// Time O(n), space O(1).
func MoveZeros(a []int) {
	j := 0 // next position for a non-zero value
	for i := range a {
		if a[i] != 0 {
			a[i], a[j] = a[j], a[i]
			j++
		}
	}
}
