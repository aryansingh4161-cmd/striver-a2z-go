package shortestpalindrome

import (
	"math/rand"
	"testing"
)

func TestShortest(t *testing.T) {
	cases := []struct{ s, want string }{{"aacecaaa", "aaacecaaa"}, {"abcd", "dcbabcd"}, {"", ""}, {"aba", "aba"}, {"ab", "bab"}}
	for _, c := range cases {
		if got := Shortest(c.s); got != c.want {
			t.Errorf("%q got %q", c.s, got)
		}
	}
	if MinCharsToAdd("abc") != 2 || MinCharsToAdd("aacecaaa") != 1 {
		t.Error("counts")
	}
	r := rand.New(rand.NewSource(4))
	for it := 0; it < 500; it++ {
		b := make([]byte, r.Intn(15))
		for i := range b {
			b[i] = "abc"[r.Intn(3)]
		}
		if Shortest(string(b)) != Brute(string(b)) {
			t.Fatalf("%q", b)
		}
	}
}
