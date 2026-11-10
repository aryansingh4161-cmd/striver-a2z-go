package assigncookies

import "sort"

// FindContentChildren returns how many children can be satisfied. O(n log n).
func FindContentChildren(greed, sizes []int) int {
	g := append([]int{}, greed...)
	s := append([]int{}, sizes...)
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			i++
		}
		j++
	}
	return i
}
