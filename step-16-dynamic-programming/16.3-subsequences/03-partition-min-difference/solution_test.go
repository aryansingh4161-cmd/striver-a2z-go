package partitionmindifference

import (
	"math/rand"
	"testing"
)

func brute(a []int) int {
	total := 0
	for _, x := range a {
		total += x
	}
	best := total
	for mask := 0; mask < 1<<len(a); mask++ {
		s := 0
		for i, x := range a {
			if mask>>i&1 == 1 {
				s += x
			}
		}
		d := total - 2*s
		if d < 0 {
			d = -d
		}
		best = min(best, d)
	}
	return best
}

func TestMinDifference(t *testing.T) {
	cases := []struct {
		a    []int
		want int
	}{{[]int{1, 6, 11, 5}, 1}, {[]int{1, 2, 7}, 4}, {[]int{3, 9, 7, 3}, 2}, {[]int{8}, 8}}
	for _, c := range cases {
		if got := MinDifference(c.a); got != c.want {
			t.Errorf("%v got %d want %d", c.a, got, c.want)
		}
	}
	r := rand.New(rand.NewSource(3))
	for it := 0; it < 300; it++ {
		a := make([]int, 1+r.Intn(10))
		for i := range a {
			a[i] = 1 + r.Intn(15)
		}
		if got, want := MinDifference(a), brute(a); got != want {
			t.Fatalf("%v: %d vs %d", a, got, want)
		}
	}
}
