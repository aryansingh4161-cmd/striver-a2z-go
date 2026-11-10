package subarrayskdifferentintegers

// SubarraysWithKDistinct counts subarrays with exactly k distinct values. O(n).
func SubarraysWithKDistinct(a []int, k int) int {
	return atMost(a, k) - atMost(a, k-1)
}

func atMost(a []int, k int) int {
	if k < 0 {
		return 0
	}
	count := map[int]int{}
	l, total := 0, 0
	for r, x := range a {
		count[x]++
		for len(count) > k {
			count[a[l]]--
			if count[a[l]] == 0 {
				delete(count, a[l])
			}
			l++
		}
		total += r - l + 1
	}
	return total
}

// SubarraysWithKDistinctBrute checks all subarrays. O(n^2).
func SubarraysWithKDistinctBrute(a []int, k int) int {
	total := 0
	for i := range a {
		seen := map[int]bool{}
		for j := i; j < len(a); j++ {
			seen[a[j]] = true
			if len(seen) == k {
				total++
			} else if len(seen) > k {
				break
			}
		}
	}
	return total
}
