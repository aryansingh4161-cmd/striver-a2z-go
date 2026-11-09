package search2dmatrix

// SearchMatrix treats the matrix as one sorted array. Time O(log(n*m)).
func SearchMatrix(mat [][]int, target int) bool {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return false
	}
	n, m := len(mat), len(mat[0])
	lo, hi := 0, n*m-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		v := mat[mid/m][mid%m]
		switch {
		case v == target:
			return true
		case v < target:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return false
}
