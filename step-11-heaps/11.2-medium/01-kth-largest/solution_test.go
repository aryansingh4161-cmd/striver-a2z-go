package kthlargest

import (
	"math/rand"
	"slices"
	"testing"
)

func TestKthLargest(t *testing.T) {
	if KthLargestHeap([]int{3, 2, 1, 5, 6, 4}, 2) != 5 || KthLargestQuickselect([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4) != 4 {
		t.Error("examples")
	}
	r := rand.New(rand.NewSource(3))
	for it := 0; it < 200; it++ {
		a := make([]int, r.Intn(20)+1)
		for i := range a {
			a[i] = r.Intn(15)
		}
		k := r.Intn(len(a)) + 1
		s := append([]int{}, a...)
		slices.Sort(s)
		want := s[len(s)-k]
		if KthLargestHeap(a, k) != want || KthLargestQuickselect(a, k) != want {
			t.Fatalf("%v k=%d", a, k)
		}
	}
}
