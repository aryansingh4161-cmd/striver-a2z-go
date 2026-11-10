package countsquares

// CountSquares sums, for every cell, the largest all-ones square ending there. O(n·m).
func CountSquares(mat [][]int) int {
	n, m := len(mat), len(mat[0])
	dp := make([][]int, n)
	total := 0
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			switch {
			case mat[i][j] == 0:
				dp[i][j] = 0
			case i == 0 || j == 0:
				dp[i][j] = 1
			default:
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
			total += dp[i][j]
		}
	}
	return total
}

// Brute checks every square. For tests.
func Brute(mat [][]int) int {
	n, m, count := len(mat), len(mat[0]), 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for s := 1; i+s <= n && j+s <= m; s++ {
				ok := true
				for r := i; r < i+s && ok; r++ {
					for c := j; c < j+s; c++ {
						if mat[r][c] == 0 {
							ok = false
							break
						}
					}
				}
				if !ok {
					break // a bigger square would contain this zero too
				}
				count++
			}
		}
	}
	return count
}
