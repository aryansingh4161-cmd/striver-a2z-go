package nqueens

import "strings"

// SolveNQueens returns every valid board. Time O(n!), space O(n).
func SolveNQueens(n int) [][]string {
	res := [][]string{}
	queenCol := make([]int, n) // queenCol[r] = column of the queen in row r
	cols := make([]bool, n)
	diag1 := make([]bool, 2*n-1) // r + c
	diag2 := make([]bool, 2*n-1) // r - c + n - 1
	var place func(r int)
	place = func(r int) {
		if r == n {
			res = append(res, render(queenCol))
			return
		}
		for c := 0; c < n; c++ {
			if cols[c] || diag1[r+c] || diag2[r-c+n-1] {
				continue
			}
			cols[c], diag1[r+c], diag2[r-c+n-1] = true, true, true
			queenCol[r] = c
			place(r + 1)
			cols[c], diag1[r+c], diag2[r-c+n-1] = false, false, false
		}
	}
	place(0)
	return res
}

func render(queenCol []int) []string {
	n := len(queenCol)
	board := make([]string, n)
	for r, c := range queenCol {
		board[r] = strings.Repeat(".", c) + "Q" + strings.Repeat(".", n-c-1)
	}
	return board
}
