package foursum

import "sort"

// FourSum returns all unique quadruplets summing to target, sorted.
// Time O(n^3), space O(1) besides the answer.
func FourSum(nums []int, target int) [][]int {
	a := append([]int{}, nums...)
	sort.Ints(a)
	n := len(a)
	res := [][]int{}
	for i := 0; i < n; i++ {
		if i > 0 && a[i] == a[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && a[j] == a[j-1] {
				continue
			}
			k, l := j+1, n-1
			for k < l {
				sum := a[i] + a[j] + a[k] + a[l]
				switch {
				case sum < target:
					k++
				case sum > target:
					l--
				default:
					res = append(res, []int{a[i], a[j], a[k], a[l]})
					k++
					l--
					for k < l && a[k] == a[k-1] {
						k++
					}
					for k < l && a[l] == a[l+1] {
						l--
					}
				}
			}
		}
	}
	return res
}
