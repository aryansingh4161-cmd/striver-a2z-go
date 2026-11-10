package nmeetings

import (
	"slices"
	"testing"
)

func TestMaxMeetings(t *testing.T) {
	if got := MaxMeetings([]int{1, 3, 0, 5, 8, 5}, []int{2, 4, 6, 7, 9, 9}); !slices.Equal(got, []int{1, 2, 4, 5}) {
		t.Errorf("got %v", got)
	}
	if got := MaxMeetings([]int{0, 1, 3}, []int{10, 2, 4}); len(got) != 2 {
		t.Errorf("got %v", got)
	}
	if got := MaxMeetings([]int{1, 2}, []int{2, 3}); len(got) != 1 { // touching ends not allowed
		t.Errorf("got %v", got)
	}
}
