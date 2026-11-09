package largestsubarrayzerosum

// LongestZeroSum returns the length of the longest subarray summing to 0.
// Time O(n), space O(n).
func LongestZeroSum(a []int) int {
	first := map[int]int{} // prefix sum -> first index
	sum, best := 0, 0
	for i, x := range a {
		sum += x
		if sum == 0 {
			best = i + 1
		} else if j, ok := first[sum]; ok {
			best = max(best, i-j)
		} else {
			first[sum] = i
		}
	}
	return best
}
