package matrixchainmultiplication

import (
	"math/rand"
	"testing"
)

func TestMCM(t *testing.T) {
	cases := []struct {
		a    []int
		want int
	}{{[]int{10, 20, 30, 40, 50}, 38000}, {[]int{10, 20, 30, 40}, 18000}, {[]int{40, 20, 30, 10, 30}, 26000}, {[]int{10, 20}, 0}}
	for _, c := range cases {
		for i, f := range []func([]int) int{Recursive, Memo, Tab} {
			if got := f(c.a); got != c.want {
				t.Errorf("f%d(%v) = %d want %d", i, c.a, got, c.want)
			}
		}
	}
	r := rand.New(rand.NewSource(1))
	for it := 0; it < 200; it++ {
		a := make([]int, 2+r.Intn(7))
		for i := range a {
			a[i] = 1 + r.Intn(20)
		}
		want := Recursive(a)
		if Memo(a) != want || Tab(a) != want {
			t.Fatalf("%v", a)
		}
	}
}
