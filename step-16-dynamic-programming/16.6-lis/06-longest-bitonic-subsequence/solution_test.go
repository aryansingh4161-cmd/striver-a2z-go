package longestbitonicsubsequence

import (
	"math/rand"
	"testing"
)

func bitonic(s []int) bool {
	i := 1
	for i < len(s) && s[i] > s[i-1] {
		i++
	}
	for i < len(s) && s[i] < s[i-1] {
		i++
	}
	return i >= len(s)
}

func brute(a []int) int {
	best := 0
	for mask := 1; mask < 1<<len(a); mask++ {
		var s []int
		for i, x := range a {
			if mask>>i&1 == 1 {
				s = append(s, x)
			}
		}
		if bitonic(s) {
			best = max(best, len(s))
		}
	}
	return best
}

func TestBitonic(t *testing.T) {
	cases := []struct {
		a    []int
		want int
	}{{[]int{1, 11, 2, 10, 4, 5, 2, 1}, 6}, {[]int{1, 2, 5, 3, 2}, 5}, {[]int{5, 4, 3}, 3}, {[]int{7}, 1}}
	for _, c := range cases {
		if got := LongestBitonic(c.a); got != c.want {
			t.Errorf("%v got %d want %d", c.a, got, c.want)
		}
	}
	r := rand.New(rand.NewSource(6))
	for it := 0; it < 300; it++ {
		a := make([]int, 1+r.Intn(10))
		for i := range a {
			a[i] = r.Intn(8)
		}
		if LongestBitonic(a) != brute(a) {
			t.Fatalf("%v", a)
		}
	}
}
