package foursum

import (
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	tests := []struct {
		a      []int
		target int
		want   [][]int
	}{
		{[]int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{[]int{2, 2, 2, 2, 2}, 8, [][]int{{2, 2, 2, 2}}},
		{[]int{1, 2, 3}, 6, [][]int{}},
	}
	for _, tc := range tests {
		if got := FourSum(tc.a, tc.target); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("FourSum(%v,%d) = %v", tc.a, tc.target, got)
		}
	}
}
