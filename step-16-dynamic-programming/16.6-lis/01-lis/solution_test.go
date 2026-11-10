package lis

import (
	"math/rand"
	"testing"
)

func TestLIS(t *testing.T) {
	fs := []func([]int) int{Recursive, Memo, TwoRows, EndingAt}
	cases := []struct {
		a    []int
		want int
	}{{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4}, {[]int{0, 1, 0, 3, 2, 3}, 4}, {[]int{7, 7, 7}, 1}, {[]int{5}, 1}}
	for _, c := range cases {
		for i, f := range fs {
			if got := f(c.a); got != c.want {
				t.Errorf("f%d(%v) = %d want %d", i, c.a, got, c.want)
			}
		}
	}
	r := rand.New(rand.NewSource(1))
	for it := 0; it < 300; it++ {
		a := make([]int, 1+r.Intn(12))
		for i := range a {
			a[i] = r.Intn(10)
		}
		want := Recursive(a)
		for i, f := range fs[1:] {
			if f(a) != want {
				t.Fatalf("f%d(%v)", i+1, a)
			}
		}
	}
}
