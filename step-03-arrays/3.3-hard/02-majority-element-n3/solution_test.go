package majorityelementn3

import (
	"slices"
	"testing"
)

func TestMajority(t *testing.T) {
	tests := []struct{ a, want []int }{
		{[]int{1, 1, 1, 3, 3, 2, 2, 2}, []int{1, 2}},
		{[]int{3, 2, 3}, []int{3}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 2, 3, 4}, []int{}},
		{[]int{2, 2}, []int{2}},
	}
	for _, tc := range tests {
		if got := Majority(tc.a); !slices.Equal(got, tc.want) {
			t.Errorf("Majority(%v) = %v", tc.a, got)
		}
		if got := MajorityMap(tc.a); !slices.Equal(got, tc.want) {
			t.Errorf("MajorityMap(%v) = %v", tc.a, got)
		}
	}
}
