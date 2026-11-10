package networkdelaytime

import "testing"

func TestNetworkDelay(t *testing.T) {
	if got := NetworkDelay([][3]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2); got != 2 {
		t.Errorf("got %d", got)
	}
	if got := NetworkDelay([][3]int{{1, 2, 1}}, 2, 2); got != -1 {
		t.Errorf("direction matters: got %d", got)
	}
	if got := NetworkDelay([][3]int{{1, 2, 5}, {1, 3, 1}, {3, 2, 1}}, 3, 1); got != 2 {
		t.Errorf("got %d", got)
	}
	if got := NetworkDelay(nil, 1, 1); got != 0 {
		t.Errorf("single node: got %d", got)
	}
}
