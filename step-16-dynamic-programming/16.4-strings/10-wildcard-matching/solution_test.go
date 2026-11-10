package wildcardmatching

import (
	"math/rand"
	"path"
	"testing"
)

func TestWildcard(t *testing.T) {
	fs := []func(string, string) bool{Recursive, Tab, Optimal}
	cases := []struct {
		s, p string
		want bool
	}{
		{"aa", "a", false}, {"aa", "*", true}, {"cb", "?a", false},
		{"adceb", "*a*b", true}, {"acdcb", "a*c?b", false}, {"", "***", true}, {"", "", true}, {"a", "", false},
	}
	for _, c := range cases {
		for i, f := range fs {
			if got := f(c.s, c.p); got != c.want {
				t.Errorf("f%d(%q,%q) = %v", i, c.s, c.p, got)
			}
		}
	}
	// Go's path.Match uses the same * and ? rules when there is no '/'.
	r := rand.New(rand.NewSource(10))
	for it := 0; it < 500; it++ {
		s, p := make([]byte, r.Intn(7)), make([]byte, r.Intn(6))
		for i := range s {
			s[i] = "ab"[r.Intn(2)]
		}
		for i := range p {
			p[i] = "ab?*"[r.Intn(4)]
		}
		want, _ := path.Match(string(p), string(s))
		for i, f := range fs {
			if f(string(s), string(p)) != want {
				t.Fatalf("f%d(%q,%q) want %v", i, s, p, want)
			}
		}
	}
}
