package numberofislands

import "testing"

func g(rows ...string) [][]byte {
	out := make([][]byte, len(rows))
	for i, r := range rows {
		out[i] = []byte(r)
	}
	return out
}

func TestNumIslands(t *testing.T) {
	if NumIslands(g("11000", "11000", "00100", "00011")) != 3 {
		t.Error("4-dir")
	}
	if NumIslands(g("101", "010", "101")) != 5 || NumIslands8(g("101", "010", "101")) != 1 {
		t.Error("diagonals")
	}
}
