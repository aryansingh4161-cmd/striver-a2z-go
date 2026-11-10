package maximumxor

import (
	"math/rand"
	"testing"
)

func TestMaxXor(t *testing.T) {
	cases := []struct {
		a    []int
		want int
	}{{[]int{3, 10, 5, 25, 2, 8}, 28}, {[]int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70}, 127}, {[]int{0}, 0}}
	for _, c := range cases {
		if FindMaximumXOR(c.a) != c.want || Brute(c.a) != c.want {
			t.Errorf("%v", c.a)
		}
	}
	r := rand.New(rand.NewSource(5))
	for it := 0; it < 300; it++ {
		a := make([]int, 1+r.Intn(20))
		for i := range a {
			a[i] = r.Intn(1 << 20)
		}
		if FindMaximumXOR(a) != Brute(a) {
			t.Fatalf("%v", a)
		}
	}
}
