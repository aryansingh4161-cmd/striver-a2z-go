package minimumplatforms

import "testing"

func TestMinPlatforms(t *testing.T) {
	tests := []struct {
		arr, dep []int
		want     int
	}{
		{[]int{900, 940, 950, 1100, 1500, 1800}, []int{910, 1200, 1120, 1130, 1900, 2000}, 3},
		{[]int{900, 1100, 1235}, []int{1000, 1200, 1240}, 1},
		{[]int{1000, 1000}, []int{1000, 1000}, 2},
		{[]int{900, 1000}, []int{1000, 1100}, 2}, // same-minute arrival and departure
	}
	for _, tc := range tests {
		if got := MinPlatforms(tc.arr, tc.dep); got != tc.want {
			t.Errorf("MinPlatforms(%v,%v) = %d", tc.arr, tc.dep, got)
		}
	}
}
