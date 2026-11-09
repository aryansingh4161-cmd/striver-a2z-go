package armstrongnumber

import "testing"

func TestIsArmstrong(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{153, true}, {370, true}, {9474, true}, {123, false}, {0, true}, {5, true}, {10, false},
	}
	for _, tc := range tests {
		if got := IsArmstrong(tc.n); got != tc.want {
			t.Errorf("IsArmstrong(%d) = %v, want %v", tc.n, got, tc.want)
		}
	}
}
