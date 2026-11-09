package gcd

import "testing"

func TestGCD(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{12, 18, 6}, {7, 13, 1}, {20, 0, 20}, {0, 5, 5}, {100, 75, 25}, {9, 9, 9},
	}
	for _, tc := range tests {
		if got := GCD(tc.a, tc.b); got != tc.want {
			t.Errorf("GCD(%d,%d) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
		if got := GCDBrute(tc.a, tc.b); got != tc.want {
			t.Errorf("GCDBrute(%d,%d) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
	if got := LCM(4, 6); got != 12 {
		t.Errorf("LCM(4,6) = %d, want 12", got)
	}
}
