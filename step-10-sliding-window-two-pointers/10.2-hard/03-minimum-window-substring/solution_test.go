package minimumwindowsubstring

import "testing"

func TestMinWindow(t *testing.T) {
	tests := []struct{ s, t, want string }{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
		{"aa", "aa", "aa"},
		{"ab", "b", "b"},
		{"bba", "ab", "ba"},
	}
	for _, tc := range tests {
		if got := MinWindow(tc.s, tc.t); got != tc.want {
			t.Errorf("MinWindow(%q,%q) = %q", tc.s, tc.t, got)
		}
	}
}
