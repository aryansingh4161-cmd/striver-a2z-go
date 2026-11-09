package searchrotatedsortedarray

import "testing"

func TestSearch(t *testing.T) {
	a := []int{4, 5, 6, 7, 0, 1, 2}
	for i, v := range a {
		if got := Search(a, v); got != i {
			t.Errorf("Search(%d) = %d, want %d", v, got, i)
		}
	}
	if got := Search(a, 3); got != -1 {
		t.Errorf("missing = %d", got)
	}
	if got := Search([]int{1}, 0); got != -1 {
		t.Errorf("single = %d", got)
	}
	if got := Search([]int{3, 1}, 1); got != 1 {
		t.Errorf("two = %d", got)
	}
}
