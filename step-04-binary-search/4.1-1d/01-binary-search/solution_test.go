package binarysearch

import "testing"

func TestSearch(t *testing.T) {
	a := []int{-1, 0, 3, 5, 9, 12}
	tests := []struct{ target, want int }{{9, 4}, {2, -1}, {-1, 0}, {12, 5}, {13, -1}}
	for _, tc := range tests {
		if got := Search(a, tc.target); got != tc.want {
			t.Errorf("Search(%d) = %d", tc.target, got)
		}
		if got := SearchRecursive(a, tc.target); got != tc.want {
			t.Errorf("SearchRecursive(%d) = %d", tc.target, got)
		}
	}
	if Search(nil, 1) != -1 {
		t.Error("empty slice")
	}
}
