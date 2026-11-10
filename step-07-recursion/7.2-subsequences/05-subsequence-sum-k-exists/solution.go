package subsequencesumkexists

// Find returns one subsequence of a with sum k, and whether it exists.
// Stops at the first answer. Time O(2^n) worst case.
func Find(a []int, k int) ([]int, bool) {
	path := []int{}
	var f func(i, sum int) bool
	f = func(i, sum int) bool {
		if i == len(a) {
			return sum == k
		}
		path = append(path, a[i])
		if f(i+1, sum+a[i]) {
			return true
		}
		path = path[:len(path)-1]
		return f(i+1, sum)
	}
	if f(0, 0) {
		return path, true
	}
	return nil, false
}
