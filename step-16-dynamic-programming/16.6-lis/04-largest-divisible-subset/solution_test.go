package largestdivisiblesubset

import (
	"math/rand"
	"slices"
	"testing"
)

func valid(s []int) bool {
	for i := range s {
		for j := range s {
			if s[i]%s[j] != 0 && s[j]%s[i] != 0 {
				return false
			}
		}
	}
	return true
}

func bruteSize(a []int) int {
	best := 0
	for mask := 0; mask < 1<<len(a); mask++ {
		var s []int
		for i, x := range a {
			if mask>>i&1 == 1 {
				s = append(s, x)
			}
		}
		if valid(s) {
			best = max(best, len(s))
		}
	}
	return best
}

func TestLargestSubset(t *testing.T) {
	if got := LargestSubset([]int{1, 2, 4, 8}); !slices.Equal(got, []int{1, 2, 4, 8}) {
		t.Errorf("got %v", got)
	}
	if got := LargestSubset([]int{1, 2, 3}); len(got) != 2 || !valid(got) {
		t.Errorf("got %v", got)
	}
	r := rand.New(rand.NewSource(4))
	for it := 0; it < 200; it++ {
		perm := r.Perm(30)
		a := make([]int, 1+r.Intn(10))
		for i := range a {
			a[i] = perm[i] + 1 // distinct positives
		}
		got := LargestSubset(a)
		if !valid(got) || len(got) != bruteSize(a) {
			t.Fatalf("%v → %v", a, got)
		}
	}
}
