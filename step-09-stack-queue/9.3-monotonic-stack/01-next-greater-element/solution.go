package nextgreaterelement

// NGEBrute scans right for each element. O(n^2).
func NGEBrute(a []int) []int {
	res := make([]int, len(a))
	for i := range a {
		res[i] = -1
		for j := i + 1; j < len(a); j++ {
			if a[j] > a[i] {
				res[i] = a[j]
				break
			}
		}
	}
	return res
}

// NGE uses a monotonic stack scanned from the right. O(n).
func NGE(a []int) []int {
	res := make([]int, len(a))
	st := []int{} // values, decreasing from bottom to top
	for i := len(a) - 1; i >= 0; i-- {
		for len(st) > 0 && st[len(st)-1] <= a[i] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			res[i] = -1
		} else {
			res[i] = st[len(st)-1]
		}
		st = append(st, a[i])
	}
	return res
}
