package mergeoverlappingintervals

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct{ in, want [][]int }{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{[][]int{{1, 10}, {2, 3}}, [][]int{{1, 10}}},
		{[][]int{{8, 9}, {1, 2}}, [][]int{{1, 2}, {8, 9}}},
		{[][]int{}, [][]int{}},
	}
	for _, tc := range tests {
		if got := Merge(tc.in); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Merge(%v) = %v", tc.in, got)
		}
	}
}
