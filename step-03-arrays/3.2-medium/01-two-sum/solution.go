package twosum

import "sort"

// TwoSum returns indices of two elements adding to target, or nil.
// Hash map. Time O(n), space O(n).
func TwoSum(a []int, target int) []int {
	seen := map[int]int{} // value -> index
	for i, x := range a {
		if j, ok := seen[target-x]; ok {
			return []int{j, i}
		}
		seen[x] = i
	}
	return nil
}

// HasPair answers yes/no with sorting + two pointers. Time O(n log n).
// It sorts a copy, so the caller's slice is unchanged.
func HasPair(a []int, target int) bool {
	b := append([]int{}, a...)
	sort.Ints(b)
	l, r := 0, len(b)-1
	for l < r {
		sum := b[l] + b[r]
		switch {
		case sum == target:
			return true
		case sum < target:
			l++
		default:
			r--
		}
	}
	return false
}
