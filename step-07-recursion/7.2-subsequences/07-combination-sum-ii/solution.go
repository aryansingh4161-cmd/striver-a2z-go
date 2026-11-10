package combinationsumii

import "sort"

// CombinationSum2 returns unique combinations (each index used once) summing to target.
func CombinationSum2(cands []int, target int) [][]int {
	a := append([]int{}, cands...)
	sort.Ints(a)
	res := [][]int{}
	path := []int{}
	var f func(start, rem int)
	f = func(start, rem int) {
		if rem == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		for j := start; j < len(a); j++ {
			if j > start && a[j] == a[j-1] {
				continue // same value already tried at this level
			}
			if a[j] > rem {
				break
			}
			path = append(path, a[j])
			f(j+1, rem-a[j])
			path = path[:len(path)-1]
		}
	}
	f(0, target)
	return res
}
