package minheapfromscratch

import (
	"math/rand"
	"slices"
	"testing"
)

func TestMinHeap(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	for it := 0; it < 100; it++ {
		a := make([]int, r.Intn(40))
		for i := range a {
			a[i] = r.Intn(50) - 25
		}
		want := append([]int{}, a...)
		slices.Sort(want)
		if got := HeapSort(a); !slices.Equal(got, want) {
			t.Fatalf("HeapSort(%v) = %v", a, got)
		}
		// push one by one
		var h MinHeap
		for _, x := range a {
			h.Push(x)
		}
		for _, w := range want {
			if x, _ := h.Pop(); x != w {
				t.Fatalf("pop order wrong")
			}
		}
		if _, ok := h.Pop(); ok {
			t.Fatal("should be empty")
		}
	}
}
