package sort012

import (
	"slices"
	"testing"
)

func TestSort012(t *testing.T) {
	cases := [][]int{{2, 0, 2, 1, 1, 0}, {2, 0, 1}, {0}, {}, {2, 2, 1, 1, 0, 0}, {1, 1, 1}}
	for _, c := range cases {
		want := append([]int{}, c...)
		slices.Sort(want)
		for name, fn := range map[string]func([]int){"count": SortCount, "dnf": SortDNF} {
			got := append([]int{}, c...)
			fn(got)
			if !slices.Equal(got, want) {
				t.Errorf("%s(%v) = %v", name, c, got)
			}
		}
	}
}
