package palindromelist

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		a    []int
		want bool
	}{{[]int{1, 2, 2, 1}, true}, {[]int{1, 2}, false}, {[]int{1, 2, 3, 2, 1}, true}, {[]int{1}, true}, {[]int{}, true}, {[]int{1, 2, 3, 1}, false}}
	for _, tc := range tests {
		head := list.FromSlice(tc.a)
		if got := IsPalindrome(head); got != tc.want {
			t.Errorf("IsPalindrome(%v) = %v", tc.a, got)
		}
		if got := list.ToSlice(head); !slices.Equal(got, tc.a) {
			t.Errorf("input not restored: %v", got)
		}
		if got := IsPalindromeStack(list.FromSlice(tc.a)); got != tc.want {
			t.Errorf("IsPalindromeStack(%v) = %v", tc.a, got)
		}
	}
}
