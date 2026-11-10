package minimummultiplications

import "testing"

func TestMinSteps(t *testing.T) {
	cases := []struct {
		arr              []int
		start, end, want int
	}{
		{[]int{2, 5, 7}, 3, 30, 2},
		{[]int{3, 4, 65}, 7, 66175, 4},
		{[]int{2}, 5, 5, 0},
		{[]int{2}, 1, 3, -1}, // powers of two never hit an odd number > 1
	}
	for _, c := range cases {
		if got := MinSteps(c.arr, c.start, c.end); got != c.want {
			t.Errorf("%v %d→%d got %d want %d", c.arr, c.start, c.end, got, c.want)
		}
	}
}
