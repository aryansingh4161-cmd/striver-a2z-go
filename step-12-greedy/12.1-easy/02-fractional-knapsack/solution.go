package fractionalknapsack

import "sort"

// MaxValue returns the best value with fractions allowed. Time O(n log n).
func MaxValue(values, weights []int, capacity int) float64 {
	idx := make([]int, len(values))
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(a, b int) bool {
		// compare v1/w1 > v2/w2 without floats: v1*w2 > v2*w1
		i, j := idx[a], idx[b]
		return values[i]*weights[j] > values[j]*weights[i]
	})
	total := 0.0
	left := capacity
	for _, i := range idx {
		if left == 0 {
			break
		}
		if weights[i] <= left {
			total += float64(values[i])
			left -= weights[i]
		} else {
			total += float64(values[i]) * float64(left) / float64(weights[i])
			left = 0
		}
	}
	return total
}
