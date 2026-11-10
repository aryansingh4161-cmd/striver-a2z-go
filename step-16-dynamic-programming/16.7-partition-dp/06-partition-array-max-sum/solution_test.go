package partitionarraymaxsum

import (
	"math/rand"
	"testing"
)

func TestPartition(t *testing.T) {
	cases := []struct {
		a       []int
		k, want int
	}{{[]int{1, 15, 7, 9, 2, 5, 10}, 3, 84}, {[]int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}, 4, 83}, {[]int{1}, 1, 1}}
	for _, c := range cases {
		if Recursive(c.a, c.k) != c.want || Tab(c.a, c.k) != c.want {
			t.Errorf("%v k=%d", c.a, c.k)
		}
	}
	r := rand.New(rand.NewSource(6))
	for it := 0; it < 300; it++ {
		a := make([]int, 1+r.Intn(10))
		for i := range a {
			a[i] = r.Intn(20)
		}
		k := 1 + r.Intn(4)
		if Recursive(a, k) != Tab(a, k) {
			t.Fatalf("%v %d", a, k)
		}
	}
}
