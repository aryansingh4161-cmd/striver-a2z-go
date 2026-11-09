package setmatrixzeroes

// SetZeroesBetter uses separate row/column flag arrays. Time O(nm), space O(n+m).
func SetZeroesBetter(m [][]int) {
	if len(m) == 0 {
		return
	}
	rows, cols := len(m), len(m[0])
	row := make([]bool, rows)
	col := make([]bool, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if m[i][j] == 0 {
				row[i], col[j] = true, true
			}
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if row[i] || col[j] {
				m[i][j] = 0
			}
		}
	}
}

// SetZeroes stores the flags in the first row and column. Time O(nm), space O(1).
func SetZeroes(m [][]int) {
	if len(m) == 0 {
		return
	}
	rows, cols := len(m), len(m[0])
	col0 := 1 // flag for column 0 (m[0][0] is used for row 0)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if m[i][j] == 0 {
				m[i][0] = 0
				if j == 0 {
					col0 = 0
				} else {
					m[0][j] = 0
				}
			}
		}
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if m[i][0] == 0 || m[0][j] == 0 {
				m[i][j] = 0
			}
		}
	}
	if m[0][0] == 0 {
		for j := 0; j < cols; j++ {
			m[0][j] = 0
		}
	}
	if col0 == 0 {
		for i := 0; i < rows; i++ {
			m[i][0] = 0
		}
	}
}
