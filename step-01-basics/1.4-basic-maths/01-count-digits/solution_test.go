package countdigits

import "testing"

func TestCountDigits(t *testing.T) {
	tests := []struct{ n, want int }{
		{12345, 5}, {7, 1}, {0, 1}, {-905, 3}, {1000000, 7},
	}
	for _, tc := range tests {
		if got := CountDigits(tc.n); got != tc.want {
			t.Errorf("CountDigits(%d) = %d, want %d", tc.n, got, tc.want)
		}
		if got := CountDigitsLog(tc.n); got != tc.want {
			t.Errorf("CountDigitsLog(%d) = %d, want %d", tc.n, got, tc.want)
		}
	}
}
