package countgoodnumbers

import "testing"

func TestCountGoodNumbers(t *testing.T) {
	tests := []struct{ n, want int }{{1, 5}, {4, 400}, {50, 564908303}}
	for _, tc := range tests {
		if got := CountGoodNumbers(tc.n); got != tc.want {
			t.Errorf("CountGoodNumbers(%d) = %d", tc.n, got)
		}
	}
}
