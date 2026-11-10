package bstiterator

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"
)

func TestIterator(t *testing.T) {
	it := Constructor(bst.Build([]int{7, 3, 15, 9, 20}))
	got := []int{}
	for it.HasNext() {
		got = append(got, it.Next())
	}
	if !slices.Equal(got, []int{3, 7, 9, 15, 20}) {
		t.Errorf("got %v", got)
	}
	empty := Constructor(nil)
	if empty.HasNext() {
		t.Error("empty")
	}
}
