package search2dmatrixii

// SearchMatrix walks from the top-right corner. Time O(n + m), space O(1).
func SearchMatrix(mat [][]int, target int) bool {
	if len(mat) == 0 {
		return false
	}
	i, j := 0, len(mat[0])-1
	for i < len(mat) && j >= 0 {
		v := mat[i][j]
		switch {
		case v == target:
			return true
		case v > target:
			j-- // everything below in this column is bigger
		default:
			i++ // everything to the left in this row is smaller
		}
	}
	return false
}
