package stockiv

import (
	"math/rand"
	"testing"
)

func TestStockIV(t *testing.T) {
	fs := []func(int, []int) int{Recursive, Optimal, TransactionNumber}
	cases := []struct {
		k    int
		p    []int
		want int
	}{{2, []int{2, 4, 1}, 2}, {2, []int{3, 2, 6, 5, 0, 3}, 7}, {1, []int{3, 2, 6, 5, 0, 3}, 4}, {0, []int{1, 9}, 0}}
	for _, c := range cases {
		for i, f := range fs {
			if got := f(c.k, c.p); got != c.want {
				t.Errorf("f%d(k=%d, %v) = %d want %d", i, c.k, c.p, got, c.want)
			}
		}
	}
	r := rand.New(rand.NewSource(4))
	for it := 0; it < 300; it++ {
		p := make([]int, 1+r.Intn(11))
		for i := range p {
			p[i] = r.Intn(20)
		}
		k := r.Intn(4)
		want := Recursive(k, p)
		if Optimal(k, p) != want || TransactionNumber(k, p) != want {
			t.Fatalf("k=%d %v", k, p)
		}
	}
}
