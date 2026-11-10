package shortestjobfirst

import "testing"

func TestAverageWait(t *testing.T) {
	if AverageWait([]int{4, 3, 7, 1, 2}) != 4 || AverageWait([]int{1, 2, 3, 4}) != 2 || AverageWait([]int{5}) != 0 {
		t.Error("examples")
	}
}
