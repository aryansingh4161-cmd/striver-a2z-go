package knapsack01

import (
	"math/rand"
	"testing"
)

func TestKnapsack(t *testing.T) {
	fs := []func([]int, []int, int) int{Recursive, Memo, Tab, Optimal}
	cases := []struct {
		wt, val []int
		W, want int
	}{
		{[]int{1, 2, 4, 5}, []int{5, 4, 8, 6}, 5, 13},
		{[]int{3, 4, 5}, []int{30, 50, 60}, 8, 90},
		{[]int{10}, []int{100}, 5, 0},
	}
	for _, c := range cases {
		for i, f := range fs {
			if got := f(c.wt, c.val, c.W); got != c.want {
				t.Errorf("f%d = %d want %d", i, got, c.want)
			}
		}
	}
	r := rand.New(rand.NewSource(6))
	for it := 0; it < 300; it++ {
		n := 1 + r.Intn(10)
		wt, val := make([]int, n), make([]int, n)
		for i := range wt {
			wt[i], val[i] = 1+r.Intn(10), r.Intn(50)
		}
		W := r.Intn(30)
		want := Recursive(wt, val, W)
		for i, f := range fs[1:] {
			if f(wt, val, W) != want {
				t.Fatalf("f%d mismatch", i+1)
			}
		}
	}
}
