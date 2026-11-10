package replacebyrank

import "sort"

// ArrayRankTransform replaces each value with its rank (1 = smallest). O(n log n).
func ArrayRankTransform(a []int) []int {
	sorted := append([]int{}, a...)
	sort.Ints(sorted)
	rank := map[int]int{}
	for _, v := range sorted {
		if _, ok := rank[v]; !ok {
			rank[v] = len(rank) + 1
		}
	}
	res := make([]int, len(a))
	for i, v := range a {
		res[i] = rank[v]
	}
	return res
}
