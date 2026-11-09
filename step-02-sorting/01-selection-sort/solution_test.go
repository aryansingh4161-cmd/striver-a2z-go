package selectionsort

import (
	"math/rand"
	"slices"
	"testing"
)

func TestSort(t *testing.T) {
	cases := [][]int{
		{13, 46, 24, 52, 20, 9},
		{5, 1, 4, 2},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{3, 1, 2, 4, 1, 5, 2, 6, 4},
		{7},
		{},
		{2, 2, 2},
		{-3, 10, 0, -7, 5},
	}
	// plus random arrays
	r := rand.New(rand.NewSource(1))
	for k := 0; k < 50; k++ {
		n := r.Intn(30)
		a := make([]int, n)
		for i := range a {
			a[i] = r.Intn(41) - 20
		}
		cases = append(cases, a)
	}
	for _, c := range cases {
		got := append([]int{}, c...)
		want := append([]int{}, c...)
		slices.Sort(want)
		Sort(got)
		if !slices.Equal(got, want) {
			t.Errorf("Sort(%v) = %v, want %v", c, got, want)
		}
	}
}
