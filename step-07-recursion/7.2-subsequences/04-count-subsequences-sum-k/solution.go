package countsubsequencessumk

// Count returns the number of subsequences of a whose sum is k. Time O(2^n).
func Count(a []int, k int) int {
	var f func(i, sum int) int
	f = func(i, sum int) int {
		if i == len(a) {
			if sum == k {
				return 1
			}
			return 0
		}
		return f(i+1, sum+a[i]) + f(i+1, sum)
	}
	return f(0, 0)
}
