package shortestjobfirst

import "sort"

// AverageWait returns the floored average waiting time under SJF. O(n log n).
func AverageWait(bt []int) int {
	if len(bt) == 0 {
		return 0
	}
	b := append([]int{}, bt...)
	sort.Ints(b)
	elapsed, totalWait := 0, 0
	for _, t := range b {
		totalWait += elapsed
		elapsed += t
	}
	return totalWait / len(b)
}
