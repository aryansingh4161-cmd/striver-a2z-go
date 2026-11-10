package subsetsii

import "sort"

// SubsetsWithDup returns all unique subsets of a (which may contain duplicates).
func SubsetsWithDup(nums []int) [][]int {
	a := append([]int{}, nums...)
	sort.Ints(a)
	res := [][]int{}
	path := []int{}
	var f func(start int)
	f = func(start int) {
		res = append(res, append([]int{}, path...))
		for j := start; j < len(a); j++ {
			if j > start && a[j] == a[j-1] {
				continue
			}
			path = append(path, a[j])
			f(j + 1)
			path = path[:len(path)-1]
		}
	}
	f(0)
	return res
}
