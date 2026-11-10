package powersetbits

// Subsets returns all subsets of a, in mask order 0..2^n-1.
func Subsets(a []int) [][]int {
	n := len(a)
	res := make([][]int, 0, 1<<n)
	for mask := 0; mask < 1<<n; mask++ {
		sub := []int{}
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				sub = append(sub, a[i])
			}
		}
		res = append(res, sub)
	}
	return res
}
