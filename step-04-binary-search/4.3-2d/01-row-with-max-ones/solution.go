package rowwithmaxones

// RowMaxOnes returns the first row with the most 1s, or -1 if there are none.
// Lower bound per row. Time O(n log m).
func RowMaxOnes(mat [][]int) int {
	best, bestRow := 0, -1
	for i, row := range mat {
		ones := len(row) - firstOne(row)
		if ones > best {
			best, bestRow = ones, i
		}
	}
	return bestRow
}

func firstOne(row []int) int {
	lo, hi, ans := 0, len(row)-1, len(row)
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if row[mid] >= 1 {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}

// RowMaxOnesStaircase walks from the top-right corner. Time O(n + m).
func RowMaxOnesStaircase(mat [][]int) int {
	if len(mat) == 0 {
		return -1
	}
	j, ans := len(mat[0])-1, -1
	for i := 0; i < len(mat) && j >= 0; i++ {
		for j >= 0 && mat[i][j] == 1 {
			j--
			ans = i
		}
	}
	return ans
}
