package pascalstriangle

import (
	"reflect"
	"slices"
	"testing"
)

func TestPascal(t *testing.T) {
	if got := Element(5, 3); got != 6 {
		t.Errorf("Element(5,3) = %d", got)
	}
	if got := Element(1, 1); got != 1 {
		t.Errorf("Element(1,1) = %d", got)
	}
	if got := Row(5); !slices.Equal(got, []int{1, 4, 6, 4, 1}) {
		t.Errorf("Row(5) = %v", got)
	}
	want := [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}}
	if got := Generate(4); !reflect.DeepEqual(got, want) {
		t.Errorf("Generate(4) = %v", got)
	}
	if got := Element(31, 16); got != 155117520 {
		t.Errorf("Element(31,16) = %d", got)
	}
}
