package setmatrixzeroes

import (
	"reflect"
	"testing"
)

func clone(m [][]int) [][]int {
	c := make([][]int, len(m))
	for i := range m {
		c[i] = append([]int{}, m[i]...)
	}
	return c
}

func TestSetZeroes(t *testing.T) {
	tests := []struct{ in, want [][]int }{
		{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{[][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
		{[][]int{{1, 0}, {1, 1}}, [][]int{{0, 0}, {1, 0}}},
		{[][]int{{1}, {0}}, [][]int{{0}, {0}}},
	}
	for _, tc := range tests {
		for name, fn := range map[string]func([][]int){"better": SetZeroesBetter, "optimal": SetZeroes} {
			m := clone(tc.in)
			fn(m)
			if !reflect.DeepEqual(m, tc.want) {
				t.Errorf("%s(%v) = %v, want %v", name, tc.in, m, tc.want)
			}
		}
	}
}
