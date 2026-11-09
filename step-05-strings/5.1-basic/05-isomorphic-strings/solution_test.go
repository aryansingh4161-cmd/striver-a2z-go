package isomorphicstrings

import "testing"

func TestIsomorphic(t *testing.T) {
	tests := []struct {
		s, t string
		want bool
	}{
		{"egg", "add", true}, {"foo", "bar", false}, {"paper", "title", true}, {"badc", "baba", false}, {"ab", "a", false},
	}
	for _, tc := range tests {
		if got := IsIsomorphic(tc.s, tc.t); got != tc.want {
			t.Errorf("IsIsomorphic(%q,%q) = %v", tc.s, tc.t, got)
		}
		if got := IsIsomorphicMaps(tc.s, tc.t); got != tc.want {
			t.Errorf("IsIsomorphicMaps(%q,%q) = %v", tc.s, tc.t, got)
		}
	}
}
