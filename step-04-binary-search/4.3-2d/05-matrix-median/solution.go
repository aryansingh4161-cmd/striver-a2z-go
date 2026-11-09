package matrixmedian

// Median returns the median of a row-sorted matrix with an odd number of cells.
// Binary search on the value. Time O(log(range) * n * log m), space O(1).
func Median(mat [][]int) int {
	n, m := len(mat), len(mat[0])
	lo, hi := mat[0][0], mat[0][m-1]
	for _, row := range mat {
		lo, hi = min(lo, row[0]), max(hi, row[m-1])
	}
	need := n * m / 2 // we need more than this many elements <= x
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if countLE(mat, mid) > need {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// countLE counts elements <= x using an upper bound per row.
func countLE(mat [][]int, x int) int {
	total := 0
	for _, row := range mat {
		lo, hi, ub := 0, len(row)-1, len(row)
		for lo <= hi {
			mid := lo + (hi-lo)/2
			if row[mid] > x {
				ub = mid
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}
		total += ub
	}
	return total
}
