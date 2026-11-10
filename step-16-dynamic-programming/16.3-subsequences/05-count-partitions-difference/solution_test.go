package countpartitionsdifference

import (
	"math/rand"
	"testing"
)

func brute(a []int, d int) int {
	c := 0
	for mask := 0; mask < 1<<len(a); mask++ {
		s1, s2 := 0, 0
		for i, x := range a {
			if mask>>i&1 == 1 {
				s1 += x
			} else {
				s2 += x
			}
		}
		if s1-s2 == d {
			c++
		}
	}
	return c
}

func TestCountPartitions(t *testing.T) {
	if got := CountPartitions([]int{5, 2, 6, 4}, 3); got != 1 {
		t.Errorf("got %d", got)
	}
	if got := CountPartitions([]int{1, 1, 1, 1}, 0); got != 6 {
		t.Errorf("got %d", got)
	}
	r := rand.New(rand.NewSource(5))
	for it := 0; it < 300; it++ {
		a := make([]int, 1+r.Intn(10))
		for i := range a {
			a[i] = r.Intn(6)
		}
		d := r.Intn(10)
		if got, want := CountPartitions(a, d), brute(a, d); got != want {
			t.Fatalf("%v d=%d: %d vs %d", a, d, got, want)
		}
	}
}
