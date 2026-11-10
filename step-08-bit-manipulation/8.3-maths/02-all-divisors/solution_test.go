package alldivisors

import (
	"slices"
	"testing"
)

func TestDivisors(t *testing.T) {
	if got := Divisors(36); !slices.Equal(got, []int{1, 2, 3, 4, 6, 9, 12, 18, 36}) {
		t.Errorf("got %v", got)
	}
	if got := Divisors(1); !slices.Equal(got, []int{1}) {
		t.Errorf("got %v", got)
	}
}
