package climbingstairs

import "testing"

func TestClimb(t *testing.T) {
	want := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	for n, w := range want {
		for name, f := range map[string]func(int) int{"rec": Recursive, "memo": Memo, "tab": Tab, "opt": Optimal} {
			if got := f(n); got != w {
				t.Errorf("%s(%d) = %d want %d", name, n, got, w)
			}
		}
	}
	if Optimal(45) != 1836311903 || Memo(45) != 1836311903 {
		t.Error("n = 45")
	}
}
