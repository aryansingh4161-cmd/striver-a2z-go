package longesthappyprefix

import (
	"math/rand"
	"testing"
)

func TestHappyPrefix(t *testing.T) {
	cases := []struct{ s, want string }{{"level", "l"}, {"ababab", "abab"}, {"abc", ""}, {"a", ""}, {"aaaa", "aaa"}}
	for _, c := range cases {
		if LongestPrefix(c.s) != c.want || Brute(c.s) != c.want {
			t.Errorf("%q", c.s)
		}
	}
	r := rand.New(rand.NewSource(5))
	for it := 0; it < 500; it++ {
		b := make([]byte, r.Intn(15))
		for i := range b {
			b[i] = "ab"[r.Intn(2)]
		}
		if LongestPrefix(string(b)) != Brute(string(b)) {
			t.Fatalf("%q", b)
		}
	}
}
