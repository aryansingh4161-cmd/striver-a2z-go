package insertinterval

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		in   [][]int
		iv   []int
		want [][]int
	}{
		{[][]int{{1, 3}, {6, 9}}, []int{2, 5}, [][]int{{1, 5}, {6, 9}}},
		{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}, [][]int{{1, 2}, {3, 10}, {12, 16}}},
		{[][]int{}, []int{5, 7}, [][]int{{5, 7}}},
		{[][]int{{1, 5}}, []int{6, 8}, [][]int{{1, 5}, {6, 8}}},
		{[][]int{{3, 5}}, []int{1, 2}, [][]int{{1, 2}, {3, 5}}},
	}
	for _, tc := range tests {
		if got := Insert(tc.in, tc.iv); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Insert(%v,%v) = %v", tc.in, tc.iv, got)
		}
	}
}
