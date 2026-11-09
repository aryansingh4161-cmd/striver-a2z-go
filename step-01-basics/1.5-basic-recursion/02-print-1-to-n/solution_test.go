package print1ton

import (
	"slices"
	"testing"
)

func TestPrint1ToN(t *testing.T) {
	want := []int{1, 2, 3, 4, 5}
	if got := CountUp(5); !slices.Equal(got, want) {
		t.Errorf("CountUp = %v", got)
	}
	if got := Backtrack(5); !slices.Equal(got, want) {
		t.Errorf("Backtrack = %v", got)
	}
}
