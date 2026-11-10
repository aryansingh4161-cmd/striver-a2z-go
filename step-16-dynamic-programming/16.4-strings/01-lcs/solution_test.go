package lcs

import (
	"math/rand"
	"testing"
)

func randStr(r *rand.Rand, n int, alpha string) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alpha[r.Intn(len(alpha))]
	}
	return string(b)
}

func TestLCS(t *testing.T) {
	fs := []func(string, string) int{Recursive, Memo, Tab, Optimal}
	cases := []struct {
		a, b string
		want int
	}{{"abcde", "ace", 3}, {"abc", "def", 0}, {"", "abc", 0}, {"abcba", "abcbcba", 5}}
	for _, c := range cases {
		for i, f := range fs {
			if got := f(c.a, c.b); got != c.want {
				t.Errorf("f%d(%q,%q) = %d want %d", i, c.a, c.b, got, c.want)
			}
		}
	}
	r := rand.New(rand.NewSource(1))
	for it := 0; it < 300; it++ {
		a, b := randStr(r, r.Intn(8), "abc"), randStr(r, r.Intn(8), "abc")
		want := Recursive(a, b)
		for i, f := range fs[1:] {
			if f(a, b) != want {
				t.Fatalf("f%d(%q,%q)", i+1, a, b)
			}
		}
	}
}
