package rotatematrix

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	in := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	want := [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}
	if got := RotateNew(in); !reflect.DeepEqual(got, want) {
		t.Errorf("RotateNew = %v", got)
	}
	Rotate(in)
	if !reflect.DeepEqual(in, want) {
		t.Errorf("Rotate = %v", in)
	}
	four := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	Rotate(four)
	want4 := [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}
	if !reflect.DeepEqual(four, want4) {
		t.Errorf("Rotate 4x4 = %v", four)
	}
}
