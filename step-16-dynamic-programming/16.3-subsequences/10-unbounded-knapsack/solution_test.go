package unboundedknapsack

import (
	"math/rand"
	"testing"
)

func TestUnbounded(t *testing.T) {
	fs := []func([]int, []int, int) int{Recursive, Memo, Optimal}
	cases := []struct {
		wt, val []int
		W, want int
	}{
		{[]int{2, 4, 6}, []int{5, 11, 13}, 10, 27},
		{[]int{1, 3, 4, 5}, []int{10, 40, 50, 70}, 8, 110},
		{[]int{5}, []int{9}, 4, 0},
	}
	for _, c := range cases {
		for i, f := range fs {
			if got := f(c.wt, c.val, c.W); got != c.want {
				t.Errorf("f%d = %d want %d", i, got, c.want)
			}
		}
	}
	r := rand.New(rand.NewSource(10))
	for it := 0; it < 200; it++ {
		n := 1 + r.Intn(5)
		wt, val := make([]int, n), make([]int, n)
		for i := range wt {
			wt[i], val[i] = 1+r.Intn(8), r.Intn(30)
		}
		W := r.Intn(25)
		want := Recursive(wt, val, W)
		if Memo(wt, val, W) != want || Optimal(wt, val, W) != want {
			t.Fatal("mismatch")
		}
	}
}
