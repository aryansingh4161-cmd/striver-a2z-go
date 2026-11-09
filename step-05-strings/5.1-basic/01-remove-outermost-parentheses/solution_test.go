package removeoutermostparentheses

import "testing"

func TestRemoveOuter(t *testing.T) {
	tests := []struct{ in, want string }{
		{"(()())(())", "()()()"},
		{"(()())(())(()(()))", "()()()()(())"},
		{"()()", ""},
		{"", ""},
	}
	for _, tc := range tests {
		if got := RemoveOuter(tc.in); got != tc.want {
			t.Errorf("RemoveOuter(%q) = %q", tc.in, got)
		}
	}
}
