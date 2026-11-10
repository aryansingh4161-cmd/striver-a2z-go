package combinationsumiii

// CombinationSum3 returns all sets of k distinct digits 1..9 summing to n.
func CombinationSum3(k, n int) [][]int {
	res := [][]int{}
	path := []int{}
	var f func(start, rem int)
	f = func(start, rem int) {
		if len(path) == k {
			if rem == 0 {
				res = append(res, append([]int{}, path...))
			}
			return
		}
		for d := start; d <= 9; d++ {
			if d > rem {
				break
			}
			path = append(path, d)
			f(d+1, rem-d)
			path = path[:len(path)-1]
		}
	}
	f(1, n)
	return res
}
