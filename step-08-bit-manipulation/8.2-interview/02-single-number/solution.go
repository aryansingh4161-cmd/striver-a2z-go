package singlenumber

// Single returns the element that appears once (others appear twice). O(n), O(1).
func Single(a []int) int {
	x := 0
	for _, v := range a {
		x ^= v
	}
	return x
}
