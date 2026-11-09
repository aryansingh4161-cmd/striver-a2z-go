package largestoddnumber

import "testing"

func TestLargestOdd(t *testing.T) {
	tests := []struct{ in, want string }{
		{"52", "5"}, {"35427", "35427"}, {"4206", ""}, {"0035", "35"}, {"7", "7"},
	}
	for _, tc := range tests {
		if got := LargestOdd(tc.in); got != tc.want {
			t.Errorf("LargestOdd(%q) = %q", tc.in, got)
		}
	}
}
