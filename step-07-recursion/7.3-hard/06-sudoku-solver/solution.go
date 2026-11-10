package sudokusolver

// Solve fills board in place. Returns false if the puzzle has no solution.
func Solve(board [][]byte) bool {
	var rowHas, colHas, boxHas [9][10]bool
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] != '.' {
				d := board[r][c] - '0'
				rowHas[r][d], colHas[c][d], boxHas[box(r, c)][d] = true, true, true
			}
		}
	}
	var solve func(pos int) bool
	solve = func(pos int) bool {
		for pos < 81 && board[pos/9][pos%9] != '.' {
			pos++ // skip filled cells
		}
		if pos == 81 {
			return true
		}
		r, c := pos/9, pos%9
		b := box(r, c)
		for d := byte(1); d <= 9; d++ {
			if rowHas[r][d] || colHas[c][d] || boxHas[b][d] {
				continue
			}
			board[r][c] = '0' + d
			rowHas[r][d], colHas[c][d], boxHas[b][d] = true, true, true
			if solve(pos + 1) {
				return true
			}
			board[r][c] = '.'
			rowHas[r][d], colHas[c][d], boxHas[b][d] = false, false, false
		}
		return false
	}
	return solve(0)
}

func box(r, c int) int { return (r/3)*3 + c/3 }
