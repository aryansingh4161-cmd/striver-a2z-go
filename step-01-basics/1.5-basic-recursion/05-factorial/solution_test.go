package factorial

import "testing"

func TestFactorial(t *testing.T) {
	for _, tc := range []struct{ n, want int }{{0, 1}, {1, 1}, {5, 120}, {10, 3628800}, {20, 2432902008176640000}} {
		if got := Factorial(tc.n); got != tc.want {
			t.Errorf("Factorial(%d) = %d, want %d", tc.n, got, tc.want)
		}
	}
}
