package surroundedregions

import "testing"

func TestSolve(t *testing.T) {
	b := [][]byte{[]byte("XXXX"), []byte("XOOX"), []byte("XXOX"), []byte("XOXX")}
	Solve(b)
	want := []string{"XXXX", "XXXX", "XXXX", "XOXX"}
	for i := range want {
		if string(b[i]) != want[i] {
			t.Errorf("row %d = %s", i, b[i])
		}
	}
}
