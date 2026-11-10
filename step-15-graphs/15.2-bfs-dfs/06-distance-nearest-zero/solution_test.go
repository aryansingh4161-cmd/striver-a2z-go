package distancenearestzero

import (
	"reflect"
	"testing"
)

func TestUpdateMatrix(t *testing.T) {
	if got := UpdateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}); !reflect.DeepEqual(got, [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}) {
		t.Errorf("got %v", got)
	}
	if got := UpdateMatrix([][]int{{1, 1, 0}}); !reflect.DeepEqual(got, [][]int{{2, 1, 0}}) {
		t.Errorf("got %v", got)
	}
}
