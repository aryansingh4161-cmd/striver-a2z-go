package maxheap

import (
	"math/rand"
	"slices"
	"testing"
)

func TestMaxHeaps(t *testing.T) {
	r := rand.New(rand.NewSource(2))
	for it := 0; it < 100; it++ {
		a := make([]int, r.Intn(30))
		for i := range a {
			a[i] = r.Intn(40)
		}
		want := append([]int{}, a...)
		slices.Sort(want)
		slices.Reverse(want)
		var h MaxHeap
		for _, x := range a {
			h.Push(x)
		}
		got := []int{}
		for h.Len() > 0 {
			got = append(got, h.Pop())
		}
		if !slices.Equal(got, want) || !slices.Equal(Drain(a), want) {
			t.Fatalf("mismatch on %v", a)
		}
	}
}
