package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{{"()[]{}", true}, {"(]", false}, {"([)]", false}, {"{[]}", true}, {"", true}, {"(", false}, {")", false}}
	for _, tc := range tests {
		if got := IsValid(tc.s); got != tc.want {
			t.Errorf("IsValid(%q) = %v", tc.s, got)
		}
	}
}
