package partitionequalsubsetsum

import subsetsum "github.com/aryansingh4161-cmd/striver-a2z-go/step-16-dynamic-programming/16.3-subsequences/01-subset-sum"

// CanPartition reports whether nums splits into two equal-sum subsets.
func CanPartition(nums []int) bool {
	total := 0
	for _, x := range nums {
		total += x
	}
	if total%2 == 1 {
		return false
	}
	return subsetsum.Optimal(nums, total/2)
}
