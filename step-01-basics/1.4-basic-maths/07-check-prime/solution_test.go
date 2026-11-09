package checkprime

import "testing"

func TestIsPrime(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{7, true}, {12, false}, {1, false}, {2, true}, {0, false}, {29, true}, {49, false}, {97, true}, {-7, false},
	}
	for _, tc := range tests {
		if got := IsPrime(tc.n); got != tc.want {
			t.Errorf("IsPrime(%d) = %v, want %v", tc.n, got, tc.want)
		}
		if got := IsPrimeBrute(tc.n); got != tc.want {
			t.Errorf("IsPrimeBrute(%d) = %v, want %v", tc.n, got, tc.want)
		}
	}
}
