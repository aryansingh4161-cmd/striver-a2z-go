package courseschedule

import "testing"

func valid(n int, pre [][2]int, order []int) bool {
	if len(order) != n {
		return false
	}
	pos := map[int]int{}
	for i, c := range order {
		pos[c] = i
	}
	for _, p := range pre {
		if pos[p[1]] > pos[p[0]] {
			return false
		}
	}
	return len(pos) == n
}

func TestCourseSchedule(t *testing.T) {
	pre := [][2]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	if o := FindOrder(4, pre); !valid(4, pre, o) {
		t.Errorf("bad order %v", o)
	}
	if CanFinish(2, [][2]int{{1, 0}, {0, 1}}) {
		t.Error("cycle: cannot finish")
	}
	if FindOrder(2, [][2]int{{1, 0}, {0, 1}}) != nil {
		t.Error("want nil")
	}
	if !CanFinish(3, nil) || len(FindOrder(3, nil)) != 3 {
		t.Error("no prereqs")
	}
}
