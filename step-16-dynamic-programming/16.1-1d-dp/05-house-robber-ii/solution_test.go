package houserobberii

import (
	"math/rand"
	"testing"
)

// brute: try every subset with no two adjacent (circularly).
func brute(a []int) int {
	n, best := len(a), 0
	for mask := 0; mask < 1<<n; mask++ {
		ok, sum := true, 0
		for i := 0; i < n; i++ {
			if mask>>i&1 == 1 {
				j := (i + 1) % n
				if j != i && mask>>j&1 == 1 {
					ok = false
				}
				sum += a[i]
			}
		}
		if ok {
			best = max(best, sum)
		}
	}
	return best
}

func TestRob(t *testing.T) {
	cases := []struct {
		a    []int
		want int
	}{{[]int{2, 3, 2}, 3}, {[]int{1, 2, 3, 1}, 4}, {[]int{1, 2, 3}, 3}, {[]int{9}, 9}, {[]int{4, 7}, 7}}
	for _, c := range cases {
		if got := Rob(c.a); got != c.want {
			t.Errorf("%v got %d want %d", c.a, got, c.want)
		}
	}
	r := rand.New(rand.NewSource(5))
	for it := 0; it < 300; it++ {
		a := make([]int, 1+r.Intn(12))
		for i := range a {
			a[i] = r.Intn(20)
		}
		if got, want := Rob(a), brute(a); got != want {
			t.Fatalf("%v: %d vs %d", a, got, want)
		}
	}
}
