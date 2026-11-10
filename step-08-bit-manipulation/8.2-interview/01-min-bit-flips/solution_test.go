package minbitflips

import "testing"

func TestMinFlips(t *testing.T) {
	for _, tc := range []struct{ a, b, want int }{{10, 7, 3}, {3, 4, 3}, {5, 5, 0}, {0, 15, 4}} {
		if got := MinFlips(tc.a, tc.b); got != tc.want {
			t.Errorf("MinFlips(%d,%d) = %d", tc.a, tc.b, got)
		}
	}
}
