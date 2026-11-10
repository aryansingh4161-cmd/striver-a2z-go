package booleanevaluation

import (
	"math/rand"
	"testing"
)

func TestBoolean(t *testing.T) {
	cases := []struct {
		s    string
		want int
	}{{"T|T&F^T", 4}, {"T^F|F", 2}, {"T", 1}, {"F", 0}, {"F|T^F", 2}}
	for _, c := range cases {
		if Recursive(c.s) != c.want || Tab(c.s) != c.want {
			t.Errorf("%q: rec %d tab %d want %d", c.s, Recursive(c.s), Tab(c.s), c.want)
		}
	}
	r := rand.New(rand.NewSource(4))
	for it := 0; it < 300; it++ {
		k := 1 + r.Intn(6)
		b := make([]byte, 2*k-1)
		for i := range b {
			if i%2 == 0 {
				b[i] = "TF"[r.Intn(2)]
			} else {
				b[i] = "&|^"[r.Intn(3)]
			}
		}
		if Recursive(string(b)) != Tab(string(b)) {
			t.Fatalf("%q", b)
		}
	}
}
