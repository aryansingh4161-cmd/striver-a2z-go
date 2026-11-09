package countsubarrayssumk

// CountBrute checks every subarray. Time O(n^2).
func CountBrute(a []int, k int) int {
	ans := 0
	for i := range a {
		sum := 0
		for j := i; j < len(a); j++ {
			sum += a[j]
			if sum == k {
				ans++
			}
		}
	}
	return ans
}

// Count uses prefix sums with a frequency map. Time O(n), space O(n).
func Count(a []int, k int) int {
	seen := map[int]int{0: 1} // prefix sum -> how many times seen
	sum, ans := 0, 0
	for _, x := range a {
		sum += x
		ans += seen[sum-k]
		seen[sum]++
	}
	return ans
}
