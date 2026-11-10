package swapnumbers

import "testing"

func TestSwapXOR(t *testing.T) {
	for _, p := range [][2]int{{5, 9}, {0, 7}, {-3, 4}, {6, 6}} {
		a, b := SwapXOR(p[0], p[1])
		if a != p[1] || b != p[0] {
			t.Errorf("SwapXOR(%d,%d) = %d,%d", p[0], p[1], a, b)
		}
	}
}
