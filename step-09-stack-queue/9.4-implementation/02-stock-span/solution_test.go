package stockspan

import (
	"slices"
	"testing"
)

func TestStockSpanner(t *testing.T) {
	var s StockSpanner
	var got []int
	for _, p := range []int{100, 80, 60, 70, 60, 75, 85} {
		got = append(got, s.Next(p))
	}
	if !slices.Equal(got, []int{1, 1, 1, 2, 1, 4, 6}) {
		t.Errorf("got %v", got)
	}
}
