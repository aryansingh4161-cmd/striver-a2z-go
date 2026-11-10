package stockfee

import (
	"math/rand"
	"testing"
)

func TestFee(t *testing.T) {
	cases := []struct {
		p         []int
		fee, want int
	}{{[]int{1, 3, 2, 8, 4, 9}, 2, 8}, {[]int{1, 3, 7, 5, 10, 3}, 3, 6}, {[]int{5, 4}, 1, 0}}
	for _, c := range cases {
		if Recursive(c.p, c.fee) != c.want || Optimal(c.p, c.fee) != c.want {
			t.Errorf("%v fee %d", c.p, c.fee)
		}
	}
	r := rand.New(rand.NewSource(6))
	for it := 0; it < 300; it++ {
		p := make([]int, 1+r.Intn(12))
		for i := range p {
			p[i] = r.Intn(20)
		}
		fee := r.Intn(5)
		if Recursive(p, fee) != Optimal(p, fee) {
			t.Fatalf("%v %d", p, fee)
		}
	}
}
