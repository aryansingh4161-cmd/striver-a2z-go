package identicaltrees

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestIsSame(t *testing.T) {
	N := tree.Null
	f := tree.FromLevel
	if !IsSame(f([]int{1, 2, 3}), f([]int{1, 2, 3})) {
		t.Error("same")
	}
	if IsSame(f([]int{1, 2}), f([]int{1, N, 2})) {
		t.Error("different shape")
	}
	if IsSame(f([]int{1, 2, 1}), f([]int{1, 1, 2})) {
		t.Error("different values")
	}
	if !IsSame(nil, nil) {
		t.Error("nil")
	}
}
