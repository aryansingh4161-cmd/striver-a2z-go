package nextsmallerelement

import (
	"slices"
	"testing"
)

func TestNSE(t *testing.T) {
	if got := NSE([]int{4, 8, 5, 2, 25}); !slices.Equal(got, []int{2, 5, 2, -1, -1}) {
		t.Errorf("NSE = %v", got)
	}
	if got := PSE([]int{4, 5, 2, 10, 8}); !slices.Equal(got, []int{-1, 4, -1, 2, 2}) {
		t.Errorf("PSE = %v", got)
	}
}
