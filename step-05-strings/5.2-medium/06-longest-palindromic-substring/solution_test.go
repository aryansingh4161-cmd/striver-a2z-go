package longestpalindromicsubstring

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct{ in, want string }{
		{"babad", "bab"}, {"cbbd", "bb"}, {"a", "a"}, {"forgeeksskeegfor", "geeksskeeg"}, {"", ""}, {"ac", "a"},
	}
	for _, tc := range tests {
		if got := LongestPalindrome(tc.in); got != tc.want {
			t.Errorf("LongestPalindrome(%q) = %q", tc.in, got)
		}
	}
}
