package subsetsum

import (
	"math/rand"
	"testing"
)

func brute(arr []int, target int) bool {
	for mask := 0; mask < 1<<len(arr); mask++ {
		s := 0
		for i, x := range arr {
			if mask>>i&1 == 1 {
				s += x
			}
		}
		if s == target {
			return true
		}
	}
	return false
}

func TestSubsetSum(t *testing.T) {
	fs := []func([]int, int) bool{Recursive, Memo, Tab, Optimal}
	cases := []struct {
		a    []int
		t    int
		want bool
	}{
		{[]int{1, 2, 3, 4}, 4, true},
		{[]int{4, 3, 2, 1}, 11, false},
		{[]int{1, 2, 3}, 5, true},
		{[]int{6}, 3, false},
	}
	for _, c := range cases {
		for i, f := range fs {
			if got := f(c.a, c.t); got != c.want {
				t.Errorf("f%d(%v, %d) = %v", i, c.a, c.t, got)
			}
		}
	}
	r := rand.New(rand.NewSource(1))
	for it := 0; it < 300; it++ {
		a := make([]int, 1+r.Intn(10))
		for i := range a {
			a[i] = 1 + r.Intn(10)
		}
		target := 1 + r.Intn(40)
		want := brute(a, target)
		for i, f := range fs {
			if f(a, target) != want {
				t.Fatalf("f%d(%v, %d) want %v", i, a, target, want)
			}
		}
	}
}
