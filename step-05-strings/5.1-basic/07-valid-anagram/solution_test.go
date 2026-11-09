package validanagram

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s, t string
		want bool
	}{{"anagram", "nagaram", true}, {"rat", "car", false}, {"a", "ab", false}, {"", "", true}}
	for _, tc := range tests {
		if got := IsAnagram(tc.s, tc.t); got != tc.want {
			t.Errorf("IsAnagram(%q,%q) = %v", tc.s, tc.t, got)
		}
		if got := IsAnagramSort(tc.s, tc.t); got != tc.want {
			t.Errorf("IsAnagramSort(%q,%q) = %v", tc.s, tc.t, got)
		}
	}
}
