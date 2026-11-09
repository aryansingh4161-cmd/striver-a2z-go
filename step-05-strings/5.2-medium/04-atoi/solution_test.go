package atoi

import "testing"

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{"42", 42}, {"   -42", -42}, {"4193 with words", 4193}, {"words 987", 0},
		{"-91283472332", -2147483648}, {"91283472332", 2147483647}, {"+1", 1}, {"", 0}, {"-", 0}, {"00012", 12},
	}
	for _, tc := range tests {
		if got := MyAtoi(tc.in); got != tc.want {
			t.Errorf("MyAtoi(%q) = %d", tc.in, got)
		}
	}
}
