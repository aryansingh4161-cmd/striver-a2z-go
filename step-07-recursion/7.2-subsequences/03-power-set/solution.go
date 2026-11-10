package powerset

// SubsetsRecursive uses pick / not pick.
func SubsetsRecursive(a []int) [][]int {
	res := [][]int{}
	path := []int{}
	var f func(i int)
	f = func(i int) {
		if i == len(a) {
			res = append(res, append([]int{}, path...))
			return
		}
		path = append(path, a[i])
		f(i + 1)
		path = path[:len(path)-1]
		f(i + 1)
	}
	f(0)
	return res
}

// SubsetsBits uses every bitmask from 0 to 2^n - 1.
func SubsetsBits(a []int) [][]int {
	n := len(a)
	res := make([][]int, 0, 1<<n)
	for mask := 0; mask < 1<<n; mask++ {
		sub := []int{}
		for i := 0; i < n; i++ {
			if mask>>i&1 == 1 {
				sub = append(sub, a[i])
			}
		}
		res = append(res, sub)
	}
	return res
}
