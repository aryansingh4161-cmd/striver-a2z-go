package majorityelement

import "testing"

func TestMajority(t *testing.T) {
	tests := []struct {
		a    []int
		want int
	}{
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{3, 2, 3}, 3},
		{[]int{7, 7, 5, 7, 5, 1, 5, 7, 5, 5, 7, 7, 5, 5, 5, 5}, 5},
		{[]int{1, 2, 3}, -1},
		{[]int{4}, 4},
	}
	for _, tc := range tests {
		if got := Majority(tc.a); got != tc.want {
			t.Errorf("Majority(%v) = %d", tc.a, got)
		}
		if got := MajorityMap(tc.a); got != tc.want {
			t.Errorf("MajorityMap(%v) = %d", tc.a, got)
		}
	}
}
