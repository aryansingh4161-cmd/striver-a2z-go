package leftrotatebyd

import (
	"slices"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		a    []int
		d    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{4, 5, 6, 7, 1, 2, 3}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 4, []int{2, 3, 1}},
		{[]int{1, 2}, 0, []int{1, 2}},
	}
	for _, tc := range tests {
		for name, fn := range map[string]func([]int, int){"temp": RotateLeftTemp, "reversal": RotateLeft} {
			a := append([]int{}, tc.a...)
			fn(a, tc.d)
			if !slices.Equal(a, tc.want) {
				t.Errorf("%s(%v, %d) = %v, want %v", name, tc.a, tc.d, a, tc.want)
			}
		}
	}
	r := []int{1, 2, 3, 4, 5, 6, 7}
	RotateRight(r, 3)
	if !slices.Equal(r, []int{5, 6, 7, 1, 2, 3, 4}) {
		t.Errorf("RotateRight = %v", r)
	}
}
