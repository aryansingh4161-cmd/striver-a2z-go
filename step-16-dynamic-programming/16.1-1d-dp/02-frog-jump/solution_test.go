package frogjump

import (
	"math/rand"
	"testing"
)

func TestFrog(t *testing.T) {
	fs := []func([]int) int{Recursive, Memo, Tab, Optimal}
	cases := []struct {
		h    []int
		want int
	}{
		{[]int{10, 20, 30, 10}, 20},
		{[]int{10, 50, 10}, 0},
		{[]int{7}, 0},
		{[]int{30, 10, 60, 10, 60, 50}, 40},
	}
	for _, c := range cases {
		for i, f := range fs {
			if got := f(c.h); got != c.want {
				t.Errorf("f%d(%v) = %d want %d", i, c.h, got, c.want)
			}
		}
	}
	r := rand.New(rand.NewSource(1))
	for it := 0; it < 200; it++ {
		h := make([]int, 1+r.Intn(15))
		for i := range h {
			h[i] = r.Intn(100)
		}
		want := Recursive(h)
		for i, f := range fs[1:] {
			if got := f(h); got != want {
				t.Fatalf("f%d(%v) = %d want %d", i+1, h, got, want)
			}
		}
	}
}
