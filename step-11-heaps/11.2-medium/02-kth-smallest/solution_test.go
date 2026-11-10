package kthsmallest

import (
	"math/rand"
	"slices"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	if KthSmallest([]int{7, 10, 4, 3, 20, 15}, 3) != 7 {
		t.Error("example")
	}
	r := rand.New(rand.NewSource(4))
	for it := 0; it < 200; it++ {
		a := make([]int, r.Intn(20)+1)
		for i := range a {
			a[i] = r.Intn(15) - 5
		}
		k := r.Intn(len(a)) + 1
		s := append([]int{}, a...)
		slices.Sort(s)
		if KthSmallest(a, k) != s[k-1] {
			t.Fatalf("%v k=%d", a, k)
		}
	}
}
