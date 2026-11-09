package leftrotatebyone

import (
	"slices"
	"testing"
)

func TestRotateLeftOne(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	RotateLeftOne(a)
	if !slices.Equal(a, []int{2, 3, 4, 5, 1}) {
		t.Errorf("got %v", a)
	}
	b := []int{9}
	RotateLeftOne(b)
	if b[0] != 9 {
		t.Errorf("single element changed: %v", b)
	}
	RotateLeftOne(nil)
}
