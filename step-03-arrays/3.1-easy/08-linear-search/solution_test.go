package linearsearch

import "testing"

func TestSearch(t *testing.T) {
	a := []int{6, 7, 8, 4, 1, 4}
	if got := Search(a, 4); got != 3 {
		t.Errorf("got %d", got)
	}
	if got := Search(a, 9); got != -1 {
		t.Errorf("got %d", got)
	}
	if got := Search(nil, 1); got != -1 {
		t.Errorf("got %d", got)
	}
}
