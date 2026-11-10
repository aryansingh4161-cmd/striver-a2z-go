package burstballoons

import (
	"math/rand"
	"testing"
)

// brute: try every bursting order.
func brute(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	best := 0
	for i := range nums {
		l, r := 1, 1
		if i > 0 {
			l = nums[i-1]
		}
		if i+1 < len(nums) {
			r = nums[i+1]
		}
		rest := append(append([]int(nil), nums[:i]...), nums[i+1:]...)
		best = max(best, l*nums[i]*r+brute(rest))
	}
	return best
}

func TestBurst(t *testing.T) {
	if Recursive([]int{3, 1, 5, 8}) != 167 || Tab([]int{3, 1, 5, 8}) != 167 || Tab([]int{1, 5}) != 10 {
		t.Error("examples")
	}
	r := rand.New(rand.NewSource(3))
	for it := 0; it < 200; it++ {
		a := make([]int, r.Intn(7))
		for i := range a {
			a[i] = r.Intn(10)
		}
		want := brute(a)
		if Recursive(a) != want || Tab(a) != want {
			t.Fatalf("%v", a)
		}
	}
}
