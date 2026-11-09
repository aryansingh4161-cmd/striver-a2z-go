package palindromenumber

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{121, true}, {123, false}, {-121, false}, {0, true}, {10, false}, {1221, true},
	}
	for _, tc := range tests {
		if got := IsPalindrome(tc.n); got != tc.want {
			t.Errorf("IsPalindrome(%d) = %v, want %v", tc.n, got, tc.want)
		}
	}
}
