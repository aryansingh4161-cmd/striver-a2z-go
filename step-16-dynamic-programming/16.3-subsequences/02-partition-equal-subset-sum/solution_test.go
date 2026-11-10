package partitionequalsubsetsum

import "testing"

func TestCanPartition(t *testing.T) {
	cases := []struct {
		a    []int
		want bool
	}{
		{[]int{1, 5, 11, 5}, true},
		{[]int{1, 2, 3, 5}, false},
		{[]int{2, 2}, true},
		{[]int{1}, false},
		{[]int{3, 3, 3, 4, 5}, true},
	}
	for _, c := range cases {
		if got := CanPartition(c.a); got != c.want {
			t.Errorf("%v got %v", c.a, got)
		}
	}
}
