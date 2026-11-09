package nextpermutation

// NextPermutation rearranges a into the next permutation in place.
// Time O(n), space O(1).
func NextPermutation(a []int) {
	n := len(a)
	// 1. find the break point
	i := n - 2
	for i >= 0 && a[i] >= a[i+1] {
		i--
	}
	if i >= 0 {
		// 2. swap with the smallest bigger element on the right
		j := n - 1
		for a[j] <= a[i] {
			j--
		}
		a[i], a[j] = a[j], a[i]
	}
	// 3. reverse the suffix (whole array if no break point)
	for l, r := i+1, n-1; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}
}
