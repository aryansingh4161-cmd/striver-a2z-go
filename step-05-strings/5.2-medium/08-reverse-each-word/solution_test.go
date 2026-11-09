package reverseeachword

import "testing"

func TestReverseEachWord(t *testing.T) {
	tests := []struct{ in, want string }{
		{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
		{"Mr Ding", "rM gniD"},
		{"  ab  cd ", "  ba  dc "},
		{"", ""},
	}
	for _, tc := range tests {
		if got := ReverseEachWord(tc.in); got != tc.want {
			t.Errorf("ReverseEachWord(%q) = %q", tc.in, got)
		}
	}
}
