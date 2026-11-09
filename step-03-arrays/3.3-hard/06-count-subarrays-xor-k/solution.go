package countsubarraysxork

// CountXOR counts subarrays whose XOR equals k. Time O(n), space O(n).
func CountXOR(a []int, k int) int {
	seen := map[int]int{0: 1} // prefix xor -> count
	xr, ans := 0, 0
	for _, x := range a {
		xr ^= x
		ans += seen[xr^k]
		seen[xr]++
	}
	return ans
}
