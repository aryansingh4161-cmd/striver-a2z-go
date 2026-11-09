package sortlist

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"
)

func TestSort(t *testing.T) {
	cases := [][]int{{4, 2, 1, 3}, {-1, 5, 3, 4, 0}, {}, {1}, {2, 2, 1}}
	r := rand.New(rand.NewSource(6))
	for k := 0; k < 30; k++ {
		a := make([]int, r.Intn(25))
		for i := range a {
			a[i] = r.Intn(20) - 10
		}
		cases = append(cases, a)
	}
	for _, a := range cases {
		want := append([]int{}, a...)
		slices.Sort(want)
		if got := list.ToSlice(Sort(list.FromSlice(a))); !slices.Equal(got, want) {
			t.Errorf("Sort(%v) = %v", a, got)
		}
		if got := list.ToSlice(SortCopy(list.FromSlice(a))); !slices.Equal(got, want) {
			t.Errorf("SortCopy(%v) = %v", a, got)
		}
	}
}
