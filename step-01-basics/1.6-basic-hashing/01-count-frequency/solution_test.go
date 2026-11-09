package countfrequency

import (
	"slices"
	"testing"
)

func TestFrequency(t *testing.T) {
	a := []int{10, 5, 10, 15, 10, 5}
	want := []Pair{{10, 3}, {5, 2}, {15, 1}}
	if got := Frequency(a); !slices.Equal(got, want) {
		t.Errorf("Frequency = %v", got)
	}
	if got := FrequencyBrute(a); !slices.Equal(got, want) {
		t.Errorf("FrequencyBrute = %v", got)
	}
	if got := Frequency(nil); len(got) != 0 {
		t.Errorf("empty input: %v", got)
	}
}
