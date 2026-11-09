package reversenumber

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct{ n, want int }{
		{123, 321}, {1200, 21}, {-45, -54}, {0, 0}, {7, 7},
	}
	for _, tc := range tests {
		if got := Reverse(tc.n); got != tc.want {
			t.Errorf("Reverse(%d) = %d, want %d", tc.n, got, tc.want)
		}
	}
	if got := Reverse32(1534236469); got != 0 {
		t.Errorf("Reverse32 overflow = %d, want 0", got)
	}
}
