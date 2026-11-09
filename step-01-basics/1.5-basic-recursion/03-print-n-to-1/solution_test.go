package printnto1

import (
	"slices"
	"testing"
)

func TestPrintNTo1(t *testing.T) {
	want := []int{4, 3, 2, 1}
	if got := CountDown(4); !slices.Equal(got, want) {
		t.Errorf("CountDown = %v", got)
	}
	if got := Backtrack(4); !slices.Equal(got, want) {
		t.Errorf("Backtrack = %v", got)
	}
}
