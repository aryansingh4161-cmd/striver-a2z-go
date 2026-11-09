package unionofsortedarrays

import (
	"slices"
	"testing"
)

func TestUnion(t *testing.T) {
	if got := Union([]int{1, 1, 2, 3, 4, 5}, []int{2, 3, 4, 4, 5, 6}); !slices.Equal(got, []int{1, 2, 3, 4, 5, 6}) {
		t.Errorf("Union = %v", got)
	}
	if got := Union(nil, []int{1, 1, 2}); !slices.Equal(got, []int{1, 2}) {
		t.Errorf("Union with empty = %v", got)
	}
}

func TestIntersection(t *testing.T) {
	if got := Intersection([]int{1, 2, 2, 3, 3, 4, 5, 6}, []int{2, 3, 3, 5, 6, 6, 7}); !slices.Equal(got, []int{2, 3, 3, 5, 6}) {
		t.Errorf("Intersection = %v", got)
	}
	if got := Intersection([]int{1, 2}, []int{3, 4}); len(got) != 0 {
		t.Errorf("no common = %v", got)
	}
}
