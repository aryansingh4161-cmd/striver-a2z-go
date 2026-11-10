package sudokusolver

import "testing"

func TestSolve(t *testing.T) {
	puzzle := []string{
		"53..7....",
		"6..195...",
		".98....6.",
		"8...6...3",
		"4..8.3..1",
		"7...2...6",
		".6....28.",
		"...419..5",
		"....8..79",
	}
	want := []string{
		"534678912",
		"672195348",
		"198342567",
		"859761423",
		"426853791",
		"713924856",
		"961537284",
		"287419635",
		"345286179",
	}
	board := make([][]byte, 9)
	for i, row := range puzzle {
		board[i] = []byte(row)
	}
	if !Solve(board) {
		t.Fatal("no solution found")
	}
	for i := range want {
		if string(board[i]) != want[i] {
			t.Errorf("row %d = %s, want %s", i, board[i], want[i])
		}
	}
}
