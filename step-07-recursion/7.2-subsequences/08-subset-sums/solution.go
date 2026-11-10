package subsetsums

import "sort"

// SubsetSums returns the sorted sums of all 2^n subsets.
func SubsetSums(a []int) []int {
	res := make([]int, 0, 1<<len(a))
	var f func(i, sum int)
	f = func(i, sum int) {
		if i == len(a) {
			res = append(res, sum)
			return
		}
		f(i+1, sum+a[i])
		f(i+1, sum)
	}
	f(0, 0)
	sort.Ints(res)
	return res
}
