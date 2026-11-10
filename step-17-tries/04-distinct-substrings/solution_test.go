package distinctsubstrings

import (
	"math/rand"
	"testing"
)

func TestDistinct(t *testing.T) {
	cases := []struct {
		s    string
		want int
	}{{"abab", 7}, {"aaa", 3}, {"abc", 6}, {"", 0}}
	for _, c := range cases {
		if CountDistinct(c.s) != c.want || Brute(c.s) != c.want {
			t.Errorf("%q", c.s)
		}
	}
	r := rand.New(rand.NewSource(4))
	for it := 0; it < 300; it++ {
		b := make([]byte, r.Intn(15))
		for i := range b {
			b[i] = "abc"[r.Intn(3)]
		}
		if CountDistinct(string(b)) != Brute(string(b)) {
			t.Fatalf("%q", b)
		}
	}
}
