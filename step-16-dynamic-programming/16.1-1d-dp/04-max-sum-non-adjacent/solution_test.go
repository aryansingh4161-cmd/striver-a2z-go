package maxsumnonadjacent

import (
	"math/rand"
	"testing"
)

func TestMaxSum(t *testing.T) {
	fs := []func([]int) int{Recursive, Memo, Tab, Optimal}
	cases := []struct {
		a    []int
		want int
	}{
		{[]int{2, 1, 4, 9}, 11},
		{[]int{1, 2, 4}, 5},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{5}, 5},
	}
	for _, c := range cases {
		for i, f := range fs {
			if got := f(c.a); got != c.want {
				t.Errorf("f%d(%v) = %d want %d", i, c.a, got, c.want)
			}
		}
	}
	r := rand.New(rand.NewSource(4))
	for it := 0; it < 200; it++ {
		a := make([]int, 1+r.Intn(15))
		for i := range a {
			a[i] = r.Intn(30)
		}
		want := Recursive(a)
		for i, f := range fs[1:] {
			if f(a) != want {
				t.Fatalf("f%d(%v)", i+1, a)
			}
		}
	}
}
