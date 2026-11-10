package countpartitionsdifference

import countsubsetssumk "github.com/aryansingh4161-cmd/striver-a2z-go/step-16-dynamic-programming/16.3-subsequences/04-count-subsets-sum-k"

// CountPartitions counts splits with sum(S1) − sum(S2) = d.
func CountPartitions(arr []int, d int) int {
	total := 0
	for _, x := range arr {
		total += x
	}
	if total-d < 0 || (total-d)%2 == 1 {
		return 0
	}
	return countsubsetssumk.Optimal(arr, (total-d)/2)
}
