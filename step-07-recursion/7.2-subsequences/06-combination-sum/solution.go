package combinationsum

// CombinationSum returns all combinations (numbers reusable) that add up to target.
func CombinationSum(cands []int, target int) [][]int {
	res := [][]int{}
	path := []int{}
	var f func(i, rem int)
	f = func(i, rem int) {
		if i == len(cands) {
			if rem == 0 {
				res = append(res, append([]int{}, path...))
			}
			return
		}
		if cands[i] <= rem {
			path = append(path, cands[i])
			f(i, rem-cands[i]) // stay: can use it again
			path = path[:len(path)-1]
		}
		f(i+1, rem)
	}
	f(0, target)
	return res
}
