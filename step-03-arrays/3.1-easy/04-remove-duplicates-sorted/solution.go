package removeduplicatessorted

// RemoveDuplicates keeps one copy of each value at the front of the sorted slice a
// and returns how many unique values there are. Time O(n), space O(1).
func RemoveDuplicates(a []int) int {
	if len(a) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(a); j++ {
		if a[j] != a[i] {
			i++
			a[i] = a[j]
		}
	}
	return i + 1
}
