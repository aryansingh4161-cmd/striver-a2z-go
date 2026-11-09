package deletealloccurrences

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"
)

func TestDeleteAll(t *testing.T) {
	tests := []struct {
		a    []int
		key  int
		want []int
	}{
		{[]int{10, 4, 10, 6, 10}, 10, []int{4, 6}},
		{[]int{1, 1, 1}, 1, []int{}},
		{[]int{1, 2, 3}, 4, []int{1, 2, 3}},
		{[]int{2, 1, 2}, 2, []int{1}},
	}
	for _, tc := range tests {
		got, ok := list.DToSlice(DeleteAll(list.DFromSlice(tc.a), tc.key))
		if !ok || !slices.Equal(got, tc.want) {
			t.Errorf("DeleteAll(%v,%d) = %v ok=%v", tc.a, tc.key, got, ok)
		}
	}
}
