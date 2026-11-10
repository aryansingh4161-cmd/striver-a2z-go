package singlenumber

import "testing"

func TestSingle(t *testing.T) {
	if Single([]int{4, 1, 2, 1, 2}) != 4 || Single([]int{7}) != 7 || Single([]int{-1, 3, 3}) != -1 {
		t.Error("Single")
	}
}
