package maxnestingdepth

import "testing"

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{{"(1+(2*3)+((8)/4))+1", 3}, {"(1)+((2))+(((3)))", 3}, {"1+2", 0}, {"()(())", 2}}
	for _, tc := range tests {
		if got := MaxDepth(tc.in); got != tc.want {
			t.Errorf("MaxDepth(%q) = %d", tc.in, got)
		}
	}
}
