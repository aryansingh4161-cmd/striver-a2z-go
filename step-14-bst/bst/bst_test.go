package bst

import (
	"slices"
	"testing"
)

func TestBuild(t *testing.T) {
	if got := Inorder(Build([]int{5, 3, 8, 1, 4})); !slices.Equal(got, []int{1, 3, 4, 5, 8}) {
		t.Errorf("got %v", got)
	}
}
