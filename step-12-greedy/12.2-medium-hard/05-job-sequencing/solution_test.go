package jobsequencing

import "testing"

func TestSchedule(t *testing.T) {
	c, p := Schedule([]Job{{1, 4, 20}, {2, 1, 10}, {3, 1, 40}, {4, 1, 30}})
	if c != 2 || p != 60 {
		t.Errorf("got %d %d", c, p)
	}
	c, p = Schedule([]Job{{1, 2, 100}, {2, 1, 19}, {3, 2, 27}, {4, 1, 25}, {5, 1, 15}})
	if c != 2 || p != 127 {
		t.Errorf("got %d %d", c, p)
	}
}
