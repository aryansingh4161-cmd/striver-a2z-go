package threesum

import "sort"

// ThreeSum returns all unique triplets summing to 0, each sorted,
// in sorted order. Time O(n^2), space O(1) besides the answer.
func ThreeSum(nums []int) [][]int {
	a := append([]int{}, nums...)
	sort.Ints(a)
	n := len(a)
	res := [][]int{}
	for i := 0; i < n; i++ {
		if i > 0 && a[i] == a[i-1] {
			continue // same first number as before
		}
		j, k := i+1, n-1
		for j < k {
			sum := a[i] + a[j] + a[k]
			switch {
			case sum < 0:
				j++
			case sum > 0:
				k--
			default:
				res = append(res, []int{a[i], a[j], a[k]})
				j++
				k--
				for j < k && a[j] == a[j-1] {
					j++
				}
				for j < k && a[k] == a[k+1] {
					k--
				}
			}
		}
	}
	return res
}
