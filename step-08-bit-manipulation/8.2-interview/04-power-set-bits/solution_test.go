package powersetbits

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	want := [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}
	if got := Subsets([]int{1, 2, 3}); !reflect.DeepEqual(got, want) {
		t.Errorf("got %v", got)
	}
}
