package numberoflis

import (
	"math/rand"
	"testing"
)

func brute(a []int) int {
	best, count := 0, 0
	for mask := 1; mask < 1<<len(a); mask++ {
		prev, l, ok := -1<<31, 0, true
		for i, x := range a {
			if mask>>i&1 == 1 {
				if x <= prev {
					ok = false
					break
				}
				prev = x
				l++
			}
		}
		if !ok {
			continue
		}
		if l > best {
			best, count = l, 1
		} else if l == best {
			count++
		}
	}
	return count
}

func TestCountLIS(t *testing.T) {
	if CountLIS([]int{1, 3, 5, 4, 7}) != 2 || CountLIS([]int{2, 2, 2, 2, 2}) != 5 {
		t.Error("examples")
	}
	r := rand.New(rand.NewSource(7))
	for it := 0; it < 300; it++ {
		a := make([]int, 1+r.Intn(11))
		for i := range a {
			a[i] = r.Intn(6)
		}
		if got, want := CountLIS(a), brute(a); got != want {
			t.Fatalf("%v: %d vs %d", a, got, want)
		}
	}
}
