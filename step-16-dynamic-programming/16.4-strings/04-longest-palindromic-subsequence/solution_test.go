package longestpalindromicsubsequence

import (
	"math/rand"
	"testing"
)

func TestLPS(t *testing.T) {
	cases := []struct {
		s    string
		want int
	}{{"bbbab", 4}, {"bbabcbcab", 7}, {"cbbd", 2}, {"a", 1}, {"", 0}}
	for _, c := range cases {
		if LPS(c.s) != c.want || Interval(c.s) != c.want {
			t.Errorf("%q: lps %d interval %d want %d", c.s, LPS(c.s), Interval(c.s), c.want)
		}
	}
	r := rand.New(rand.NewSource(4))
	for it := 0; it < 300; it++ {
		b := make([]byte, r.Intn(12))
		for i := range b {
			b[i] = "abc"[r.Intn(3)]
		}
		if LPS(string(b)) != Interval(string(b)) {
			t.Fatalf("%q", b)
		}
	}
}
