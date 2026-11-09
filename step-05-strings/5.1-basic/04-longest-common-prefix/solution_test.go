package longestcommonprefix

import "testing"

func TestLCP(t *testing.T) {
	tests := []struct {
		in   []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"a"}, "a"},
		{[]string{"ab", "a"}, "a"},
		{[]string{}, ""},
	}
	for _, tc := range tests {
		if got := LCP(tc.in); got != tc.want {
			t.Errorf("LCP(%v) = %q", tc.in, got)
		}
		if got := LCPSort(tc.in); got != tc.want {
			t.Errorf("LCPSort(%v) = %q", tc.in, got)
		}
	}
}
