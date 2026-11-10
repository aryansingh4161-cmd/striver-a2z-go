package expressionaddoperators

import (
	"slices"
	"testing"
)

func TestAddOperators(t *testing.T) {
	tests := []struct {
		num    string
		target int
		want   []string
	}{
		{"123", 6, []string{"1*2*3", "1+2+3"}},
		{"232", 8, []string{"2*3+2", "2+3*2"}},
		{"105", 5, []string{"1*0+5", "10-5"}},
		{"3456237490", 9191, []string{}},
		{"00", 0, []string{"0*0", "0+0", "0-0"}},
	}
	for _, tc := range tests {
		got := AddOperators(tc.num, tc.target)
		slices.Sort(got)
		if !slices.Equal(got, tc.want) {
			t.Errorf("AddOperators(%q,%d) = %v", tc.num, tc.target, got)
		}
	}
}
