package rotatematrix

// RotateNew builds a rotated copy. Time O(n^2), space O(n^2).
func RotateNew(m [][]int) [][]int {
	n := len(m)
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res[j][n-1-i] = m[i][j]
		}
	}
	return res
}

// Rotate rotates the square matrix 90 degrees clockwise in place:
// transpose, then reverse every row. Time O(n^2), space O(1).
func Rotate(m [][]int) {
	n := len(m)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ { // only above the diagonal
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}
	for _, row := range m {
		for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
			row[l], row[r] = row[r], row[l]
		}
	}
}
