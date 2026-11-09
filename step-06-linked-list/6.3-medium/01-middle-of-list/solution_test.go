package middleoflist

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"
)

func TestMiddle(t *testing.T) {
	tests := []struct {
		a    []int
		want int
	}{{[]int{1, 2, 3, 4, 5}, 3}, {[]int{1, 2, 3, 4, 5, 6}, 4}, {[]int{7}, 7}, {[]int{1, 2}, 2}}
	for _, tc := range tests {
		if got := Middle(list.FromSlice(tc.a)); got.Val != tc.want {
			t.Errorf("Middle(%v) = %d", tc.a, got.Val)
		}
		if got := MiddleCount(list.FromSlice(tc.a)); got.Val != tc.want {
			t.Errorf("MiddleCount(%v) = %d", tc.a, got.Val)
		}
	}
	if Middle(nil) != nil {
		t.Error("nil list")
	}
}
