package candy

// Candy returns the minimum candies satisfying the neighbour rule.
// Two passes over one array. Time O(n), space O(n).
func Candy(r []int) int {
	n := len(r)
	c := make([]int, n)
	for i := range c {
		c[i] = 1
	}
	for i := 1; i < n; i++ {
		if r[i] > r[i-1] {
			c[i] = c[i-1] + 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if r[i] > r[i+1] {
			c[i] = max(c[i], c[i+1]+1)
		}
	}
	total := 0
	for _, x := range c {
		total += x
	}
	return total
}
