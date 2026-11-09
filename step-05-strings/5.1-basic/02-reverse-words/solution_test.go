package reversewords

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct{ in, want string }{
		{"  the sky   is blue ", "blue is sky the"},
		{"hello world", "world hello"},
		{"a", "a"},
		{"   ", ""},
	}
	for _, tc := range tests {
		if got := ReverseWords(tc.in); got != tc.want {
			t.Errorf("ReverseWords(%q) = %q", tc.in, got)
		}
		if got := ReverseWordsFields(tc.in); got != tc.want {
			t.Errorf("ReverseWordsFields(%q) = %q", tc.in, got)
		}
	}
}
