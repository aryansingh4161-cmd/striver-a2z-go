package floodfill

import (
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {
	img := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	if got := FloodFill(img, 1, 1, 2); !reflect.DeepEqual(got, [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}) {
		t.Errorf("got %v", got)
	}
	if got := FloodFill([][]int{{0, 0}, {0, 0}}, 0, 0, 0); !reflect.DeepEqual(got, [][]int{{0, 0}, {0, 0}}) {
		t.Errorf("same colour: %v", got)
	}
}
