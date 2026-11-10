package mergeintervals

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	if got := Merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}); !reflect.DeepEqual(got, [][]int{{1, 6}, {8, 10}, {15, 18}}) {
		t.Errorf("got %v", got)
	}
	if got := Merge([][]int{{1, 4}, {0, 4}}); !reflect.DeepEqual(got, [][]int{{0, 4}}) {
		t.Errorf("got %v", got)
	}
}
