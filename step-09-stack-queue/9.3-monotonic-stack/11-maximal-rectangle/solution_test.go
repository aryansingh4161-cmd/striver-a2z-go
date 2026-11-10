package maximalrectangle

import "testing"

func TestMaximalRectangle(t *testing.T) {
	m := [][]byte{[]byte("10100"), []byte("10111"), []byte("11111"), []byte("10010")}
	if got := MaximalRectangle(m); got != 6 {
		t.Errorf("got %d", got)
	}
	if MaximalRectangle([][]byte{[]byte("0")}) != 0 || MaximalRectangle([][]byte{[]byte("1")}) != 1 {
		t.Error("single cell")
	}
	if MaximalRectangle([][]byte{[]byte("111"), []byte("111")}) != 6 {
		t.Error("all ones")
	}
}
