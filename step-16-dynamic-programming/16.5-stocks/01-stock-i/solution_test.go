package stocki

import (
	"math/rand"
	"testing"
)

func TestStockI(t *testing.T) {
	cases := []struct {
		p    []int
		want int
	}{{[]int{7, 1, 5, 3, 6, 4}, 5}, {[]int{7, 6, 4, 3, 1}, 0}, {[]int{5}, 0}}
	for _, c := range cases {
		if Brute(c.p) != c.want || Optimal(c.p) != c.want {
			t.Errorf("%v", c.p)
		}
	}
	r := rand.New(rand.NewSource(1))
	for it := 0; it < 300; it++ {
		p := make([]int, 1+r.Intn(12))
		for i := range p {
			p[i] = r.Intn(20)
		}
		if Brute(p) != Optimal(p) {
			t.Fatalf("%v", p)
		}
	}
}
