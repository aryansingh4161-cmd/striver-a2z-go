package makenetworkconnected

import "testing"

func TestMakeConnected(t *testing.T) {
	cases := []struct {
		n    int
		c    [][2]int
		want int
	}{
		{4, [][2]int{{0, 1}, {0, 2}, {1, 2}}, 1},
		{6, [][2]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}, 2},
		{6, [][2]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}}, -1},
		{1, nil, 0},
	}
	for _, c := range cases {
		if got := MakeConnected(c.n, c.c); got != c.want {
			t.Errorf("n=%d got %d want %d", c.n, got, c.want)
		}
	}
}
