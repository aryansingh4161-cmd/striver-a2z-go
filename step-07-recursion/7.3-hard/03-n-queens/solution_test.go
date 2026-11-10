package nqueens

import (
	"reflect"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	want4 := [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}
	if got := SolveNQueens(4); !reflect.DeepEqual(got, want4) {
		t.Errorf("n=4: %v", got)
	}
	counts := map[int]int{1: 1, 2: 0, 3: 0, 5: 10, 6: 4, 8: 92}
	for n, c := range counts {
		if got := len(SolveNQueens(n)); got != c {
			t.Errorf("n=%d: %d solutions, want %d", n, got, c)
		}
	}
}
