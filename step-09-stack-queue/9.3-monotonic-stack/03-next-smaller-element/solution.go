package nextsmallerelement

// NSE returns the next smaller element to the right (values), -1 if none. O(n).
func NSE(a []int) []int {
	res := make([]int, len(a))
	st := []int{}
	for i := len(a) - 1; i >= 0; i-- {
		for len(st) > 0 && st[len(st)-1] >= a[i] {
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

// PSE returns the previous smaller element to the left (values), -1 if none. O(n).
func PSE(a []int) []int {
	res := make([]int, len(a))
	st := []int{}
	for i := 0; i < len(a); i++ {
		for len(st) > 0 && st[len(st)-1] >= a[i] {
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
