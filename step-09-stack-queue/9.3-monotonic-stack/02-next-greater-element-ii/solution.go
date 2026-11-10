package nextgreaterelementii

// NGECircular returns next greater elements in a circular array. O(n).
func NGECircular(a []int) []int {
	n := len(a)
	res := make([]int, n)
	st := []int{}
	for i := 2*n - 1; i >= 0; i-- {
		x := a[i%n]
		for len(st) > 0 && st[len(st)-1] <= x {
			st = st[:len(st)-1]
		}
		if i < n {
			if len(st) == 0 {
				res[i] = -1
			} else {
				res[i] = st[len(st)-1]
			}
		}
		st = append(st, x)
	}
	return res
}
