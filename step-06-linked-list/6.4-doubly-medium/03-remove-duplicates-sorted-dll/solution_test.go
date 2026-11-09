package removeduplicatessorteddll

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct{ a, want []int }{
		{[]int{1, 1, 1, 2, 3, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{5, 5}, []int{5}},
		{[]int{}, []int{}},
	}
	for _, tc := range tests {
		got, ok := list.DToSlice(RemoveDuplicates(list.DFromSlice(tc.a)))
		if !ok || !slices.Equal(got, tc.want) {
			t.Errorf("RemoveDuplicates(%v) = %v ok=%v", tc.a, got, ok)
		}
	}
}
