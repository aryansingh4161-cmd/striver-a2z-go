package rotatestring

import "testing"

func TestRotate(t *testing.T) {
	tests := []struct {
		s, g string
		want bool
	}{{"abcde", "cdeab", true}, {"abcde", "abced", false}, {"", "", true}, {"a", "ab", false}, {"aa", "aa", true}}
	for _, tc := range tests {
		if got := Rotate(tc.s, tc.g); got != tc.want {
			t.Errorf("Rotate(%q,%q) = %v", tc.s, tc.g, got)
		}
		if got := RotateBrute(tc.s, tc.g); got != tc.want {
			t.Errorf("RotateBrute(%q,%q) = %v", tc.s, tc.g, got)
		}
	}
}
