package palindromepartitioningii

import (
	"math/rand"
	"testing"
)

func TestMinCut(t *testing.T) {
	cases := []struct {
		s    string
		want int
	}{{"aab", 1}, {"a", 0}, {"ab", 1}, {"ababbbabbababa", 3}, {"racecar", 0}}
	for _, c := range cases {
		if Recursive(c.s) != c.want || Tab(c.s) != c.want {
			t.Errorf("%q: rec %d tab %d want %d", c.s, Recursive(c.s), Tab(c.s), c.want)
		}
	}
	r := rand.New(rand.NewSource(5))
	for it := 0; it < 300; it++ {
		b := make([]byte, 1+r.Intn(10))
		for i := range b {
			b[i] = "ab"[r.Intn(2)]
		}
		if Recursive(string(b)) != Tab(string(b)) {
			t.Fatalf("%q", b)
		}
	}
}
