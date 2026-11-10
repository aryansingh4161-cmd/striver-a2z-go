package tree

import (
	"slices"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	cases := [][]int{{1, 2, 3, Null, 4}, {1}, {}, {3, 9, 20, Null, Null, 15, 7}, {1, Null, 2, Null, 3}}
	for _, c := range cases {
		if got := ToLevel(FromLevel(c)); !slices.Equal(got, c) {
			t.Errorf("round trip %v → %v", c, got)
		}
	}
	r := FromLevel([]int{1, 2, 3, Null, 4})
	if Find(r, 4) == nil || Find(r, 4).Val != 4 || Find(r, 9) != nil {
		t.Error("Find")
	}
}
