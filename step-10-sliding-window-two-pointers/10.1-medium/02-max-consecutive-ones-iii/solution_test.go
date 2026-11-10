package maxconsecutiveonesiii

import (
	"math/rand"
	"testing"
)

func brute(a []int, k int) int {
	best := 0
	for i := range a {
		z := 0
		for j := i; j < len(a); j++ {
			if a[j] == 0 {
				z++
			}
			if z > k {
				break
			}
			best = max(best, j-i+1)
		}
	}
	return best
}

func TestLongestOnes(t *testing.T) {
	if LongestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2) != 6 {
		t.Error("example 1")
	}
	if LongestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3) != 10 {
		t.Error("example 2")
	}
	r := rand.New(rand.NewSource(2))
	for it := 0; it < 300; it++ {
		a := make([]int, r.Intn(20))
		for i := range a {
			a[i] = r.Intn(2)
		}
		k := r.Intn(4)
		w := brute(a, k)
		if LongestOnes(a, k) != w || LongestOnesNoShrink(a, k) != w {
			t.Fatalf("mismatch %v k=%d", a, k)
		}
	}
}
