package divideintegers

import (
	"math"
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{10, 3, 3}, {7, -3, -2}, {22, 3, 7}, {0, 5, 0}, {1, 1, 1},
		{math.MinInt32, -1, math.MaxInt32}, {math.MinInt32, 1, math.MinInt32}, {-7, -2, 3},
	}
	for _, tc := range tests {
		if got := Divide(tc.a, tc.b); got != tc.want {
			t.Errorf("Divide(%d,%d) = %d", tc.a, tc.b, got)
		}
	}
	for a := -50; a <= 50; a++ {
		for b := -7; b <= 7; b++ {
			if b != 0 && Divide(a, b) != a/b {
				t.Fatalf("Divide(%d,%d)", a, b)
			}
		}
	}
}
