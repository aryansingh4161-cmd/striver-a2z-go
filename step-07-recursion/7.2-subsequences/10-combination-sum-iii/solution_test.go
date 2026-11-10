package combinationsumiii

import (
	"reflect"
	"testing"
)

func TestCombinationSum3(t *testing.T) {
	if got := CombinationSum3(3, 7); !reflect.DeepEqual(got, [][]int{{1, 2, 4}}) {
		t.Errorf("got %v", got)
	}
	if got := CombinationSum3(3, 9); !reflect.DeepEqual(got, [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}) {
		t.Errorf("got %v", got)
	}
	if got := CombinationSum3(4, 1); len(got) != 0 {
		t.Errorf("got %v", got)
	}
}
