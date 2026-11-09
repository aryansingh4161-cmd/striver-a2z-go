package mergesortedarrays

import (
	"slices"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct{ a, b, wantA, wantB []int }{
		{[]int{1, 4, 8, 10}, []int{2, 3, 9}, []int{1, 2, 3, 4}, []int{8, 9, 10}},
		{[]int{1, 3, 5, 7}, []int{0, 2, 6, 8, 9}, []int{0, 1, 2, 3}, []int{5, 6, 7, 8, 9}},
		{[]int{5}, []int{1}, []int{1}, []int{5}},
		{[]int{1, 2}, []int{}, []int{1, 2}, []int{}},
	}
	for _, tc := range tests {
		for name, fn := range map[string]func([]int, []int){"swapsort": MergeSwapSort, "gap": MergeGap} {
			a := append([]int{}, tc.a...)
			b := append([]int{}, tc.b...)
			fn(a, b)
			if !slices.Equal(a, tc.wantA) || !slices.Equal(b, tc.wantB) {
				t.Errorf("%s(%v,%v) = %v,%v", name, tc.a, tc.b, a, b)
			}
		}
	}
	a := []int{1, 2, 3, 0, 0, 0}
	MergeIntoFirst(a, 3, []int{2, 5, 6}, 3)
	if !slices.Equal(a, []int{1, 2, 2, 3, 5, 6}) {
		t.Errorf("MergeIntoFirst = %v", a)
	}
}
