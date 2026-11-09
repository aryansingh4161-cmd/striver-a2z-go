package threesum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		a    []int
		want [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 1, 1}, [][]int{}},
		{[]int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{-2, 0, 0, 2, 2}, [][]int{{-2, 0, 2}}},
	}
	for _, tc := range tests {
		if got := ThreeSum(tc.a); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("ThreeSum(%v) = %v", tc.a, got)
		}
	}
}
