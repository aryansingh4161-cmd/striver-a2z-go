package topkfrequent

import (
	"slices"
	"testing"
)

func TestTopK(t *testing.T) {
	tests := []struct {
		a    []int
		k    int
		want []int
	}{{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}}, {[]int{1}, 1, []int{1}}, {[]int{4, 4, 4, 5, 5, 6, 6, 6, 6}, 1, []int{6}}}
	for _, tc := range tests {
		if got := TopKHeap(tc.a, tc.k); !slices.Equal(got, tc.want) {
			t.Errorf("TopKHeap(%v,%d) = %v", tc.a, tc.k, got)
		}
		if got := TopKBucket(tc.a, tc.k); !slices.Equal(got, tc.want) {
			t.Errorf("TopKBucket(%v,%d) = %v", tc.a, tc.k, got)
		}
	}
}
