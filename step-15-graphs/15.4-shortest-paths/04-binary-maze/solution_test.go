package binarymaze

import "testing"

func TestShortestPath(t *testing.T) {
	g := [][]int{{1, 1, 1, 1}, {1, 1, 0, 1}, {1, 1, 1, 1}, {1, 1, 0, 0}, {1, 0, 0, 1}}
	cases := []struct {
		src, dst [2]int
		want     int
	}{
		{[2]int{0, 1}, [2]int{2, 2}, 3},
		{[2]int{0, 0}, [2]int{0, 0}, 0},
		{[2]int{0, 0}, [2]int{4, 3}, -1}, // walled off
		{[2]int{0, 0}, [2]int{1, 2}, -1}, // dst is a wall
		{[2]int{0, 0}, [2]int{4, 0}, 4},
	}
	for _, c := range cases {
		if got := ShortestPath(g, c.src, c.dst); got != c.want {
			t.Errorf("%v→%v got %d want %d", c.src, c.dst, got, c.want)
		}
	}
}
