package fruitintobaskets

// TotalFruit returns the longest subarray with at most 2 distinct values. O(n).
func TotalFruit(f []int) int {
	count := map[int]int{}
	best, l := 0, 0
	for r, x := range f {
		count[x]++
		for len(count) > 2 {
			count[f[l]]--
			if count[f[l]] == 0 {
				delete(count, f[l])
			}
			l++
		}
		best = max(best, r-l+1)
	}
	return best
}
