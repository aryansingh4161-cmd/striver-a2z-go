package pathminimumeffort

import (
	"math/rand"
	"testing"
)

func TestMinEffort(t *testing.T) {
	cases := []struct {
		h    [][]int
		want int
	}{
		{[][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}, 2},
		{[][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}, 1},
		{[][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}}, 0},
		{[][]int{{7}}, 0},
	}
	for _, c := range cases {
		if got := MinEffort(c.h); got != c.want {
			t.Errorf("dijkstra got %d want %d", got, c.want)
		}
		if got := MinEffortBinarySearch(c.h); got != c.want {
			t.Errorf("binary search got %d want %d", got, c.want)
		}
	}
	r := rand.New(rand.NewSource(8))
	for it := 0; it < 200; it++ {
		R, C := 1+r.Intn(5), 1+r.Intn(5)
		h := make([][]int, R)
		for i := range h {
			h[i] = make([]int, C)
			for j := range h[i] {
				h[i][j] = r.Intn(20)
			}
		}
		if a, b := MinEffort(h), MinEffortBinarySearch(h); a != b {
			t.Fatalf("%v: %d vs %d", h, a, b)
		}
	}
}
