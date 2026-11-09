package longestsubarraysumk

// Longest returns the length of the longest subarray with sum k.
// Works with negative numbers. Time O(n), space O(n).
func Longest(a []int, k int) int {
	firstIndex := map[int]int{} // prefix sum -> earliest index
	sum, best := 0, 0
	for i, x := range a {
		sum += x
		if sum == k {
			best = max(best, i+1)
		}
		if j, ok := firstIndex[sum-k]; ok {
			best = max(best, i-j)
		}
		if _, ok := firstIndex[sum]; !ok {
			firstIndex[sum] = i
		}
	}
	return best
}
