package oddeven

import "testing"

func TestIsOdd(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want bool
	}{{0, false}, {1, true}, {2, false}, {7, true}, {-3, true}, {-4, false}} {
		if IsOdd(tc.n) != tc.want {
			t.Errorf("IsOdd(%d)", tc.n)
		}
	}
}
