package primsalgorithm

import "testing"

func TestMST(t *testing.T) {
	edges := [][3]int{{0, 1, 2}, {0, 3, 6}, {1, 2, 3}, {1, 3, 8}, {1, 4, 5}, {2, 4, 7}}
	sum, chosen := MST(5, edges)
	if sum != 16 || len(chosen) != 4 {
		t.Errorf("got %d %v", sum, chosen)
	}
	if s, c := MST(1, nil); s != 0 || len(c) != 0 {
		t.Error("single node")
	}
	// parallel edges: the cheaper one wins
	if s, _ := MST(2, [][3]int{{0, 1, 9}, {0, 1, 4}}); s != 4 {
		t.Errorf("got %d", s)
	}
}
