package squareroot

import "testing"

func TestSqrt(t *testing.T) {
	for _, tc := range []struct{ n, want int }{{36, 6}, {28, 5}, {0, 0}, {1, 1}, {2, 1}, {2147395599, 46339}} {
		if got := Sqrt(tc.n); got != tc.want {
			t.Errorf("Sqrt(%d) = %d", tc.n, got)
		}
		if got := SqrtLinear(tc.n); got != tc.want {
			t.Errorf("SqrtLinear(%d) = %d", tc.n, got)
		}
	}
	if got := Sqrt(1 << 62); got != 1<<31 {
		t.Errorf("big = %d", got)
	}
}
