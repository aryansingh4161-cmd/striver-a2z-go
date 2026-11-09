package findpeakelementii

// FindPeak returns [row, col] of some peak. Binary search on columns,
// using each column's maximum. Time O(n log m), space O(1).
func FindPeak(mat [][]int) [2]int {
	n, m := len(mat), len(mat[0])
	lo, hi := 0, m-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		r := 0
		for i := 1; i < n; i++ { // row of the max in column mid
			if mat[i][mid] > mat[r][mid] {
				r = i
			}
		}
		left, right := -1, -1
		if mid > 0 {
			left = mat[r][mid-1]
		}
		if mid < m-1 {
			right = mat[r][mid+1]
		}
		v := mat[r][mid]
		switch {
		case v > left && v > right:
			return [2]int{r, mid}
		case left > v:
			hi = mid - 1
		default:
			lo = mid + 1
		}
	}
	return [2]int{-1, -1}
}
