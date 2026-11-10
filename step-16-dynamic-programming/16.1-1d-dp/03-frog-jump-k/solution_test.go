package frogjumpk

import (
	"math/rand"
	"testing"
)

func TestFrogK(t *testing.T) {
	h := []int{10, 30, 40, 50, 20}
	for i, f := range []func([]int, int) int{Recursive, Memo, Tab} {
		if got := f(h, 3); got != 30 {
			t.Errorf("f%d = %d want 30", i, got)
		}
		if got := f([]int{10, 20, 10}, 1); got != 20 {
			t.Errorf("k=1 f%d = %d", i, got)
		}
	}
	r := rand.New(rand.NewSource(3))
	for it := 0; it < 200; it++ {
		h := make([]int, 1+r.Intn(10))
		for i := range h {
			h[i] = r.Intn(50)
		}
		k := 1 + r.Intn(4)
		want := Recursive(h, k)
		if Memo(h, k) != want || Tab(h, k) != want {
			t.Fatalf("%v k=%d", h, k)
		}
	}
}
