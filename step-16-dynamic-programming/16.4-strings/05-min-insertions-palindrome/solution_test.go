package mininsertionspalindrome

import (
	"math/rand"
	"testing"
)

// brute: classic interval recursion on insertions.
func brute(s string) int {
	var f func(i, j int) int
	f = func(i, j int) int {
		if i >= j {
			return 0
		}
		if s[i] == s[j] {
			return f(i+1, j-1)
		}
		return 1 + min(f(i+1, j), f(i, j-1))
	}
	return f(0, len(s)-1)
}

func TestMinInsertions(t *testing.T) {
	cases := []struct {
		s    string
		want int
	}{{"abcaa", 2}, {"mbadm", 2}, {"zzazz", 0}, {"leetcode", 5}, {"", 0}}
	for _, c := range cases {
		if got := MinInsertions(c.s); got != c.want {
			t.Errorf("%q got %d want %d", c.s, got, c.want)
		}
	}
	r := rand.New(rand.NewSource(5))
	for it := 0; it < 300; it++ {
		b := make([]byte, r.Intn(10))
		for i := range b {
			b[i] = "abc"[r.Intn(3)]
		}
		if MinInsertions(string(b)) != brute(string(b)) {
			t.Fatalf("%q", b)
		}
	}
}
