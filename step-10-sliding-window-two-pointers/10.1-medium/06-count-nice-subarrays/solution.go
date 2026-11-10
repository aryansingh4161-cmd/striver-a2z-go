package countnicesubarrays

// NumberOfSubarrays counts subarrays with exactly k odd numbers. O(n), O(1).
func NumberOfSubarrays(a []int, k int) int {
	return atMost(a, k) - atMost(a, k-1)
}

func atMost(a []int, k int) int {
	if k < 0 {
		return 0
	}
	l, odd, count := 0, 0, 0
	for r, x := range a {
		odd += x & 1
		for odd > k {
			odd -= a[l] & 1
			l++
		}
		count += r - l + 1
	}
	return count
}
