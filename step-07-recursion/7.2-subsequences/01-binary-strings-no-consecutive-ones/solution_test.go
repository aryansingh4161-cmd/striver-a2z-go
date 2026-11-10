package binarystrings

import (
	"slices"
	"testing"
)

func TestGenerate(t *testing.T) {
	if got := Generate(3); !slices.Equal(got, []string{"000", "001", "010", "100", "101"}) {
		t.Errorf("Generate(3) = %v", got)
	}
	if got := Generate(1); !slices.Equal(got, []string{"0", "1"}) {
		t.Errorf("Generate(1) = %v", got)
	}
	// counts follow Fibonacci: 2, 3, 5, 8, 13, ...
	want := []int{2, 3, 5, 8, 13, 21}
	for n := 1; n <= 6; n++ {
		if got := len(Generate(n)); got != want[n-1] {
			t.Errorf("len(Generate(%d)) = %d", n, got)
		}
	}
}
