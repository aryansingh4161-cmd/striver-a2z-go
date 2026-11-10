package distinctsubsequences

import (
	"math/rand"
	"testing"
)

func TestDistinct(t *testing.T) {
	fs := []func(string, string) int{Recursive, Tab, Optimal}
	cases := []struct {
		s, t string
		want int
	}{{"babgbag", "bag", 5}, {"rabbbit", "rabbit", 3}, {"abc", "", 1}, {"", "a", 0}, {"aaa", "aa", 3}}
	for _, c := range cases {
		for i, f := range fs {
			if got := f(c.s, c.t); got != c.want {
				t.Errorf("f%d(%q,%q) = %d want %d", i, c.s, c.t, got, c.want)
			}
		}
	}
	r := rand.New(rand.NewSource(8))
	for it := 0; it < 300; it++ {
		s, tt := make([]byte, r.Intn(10)), make([]byte, r.Intn(4))
		for i := range s {
			s[i] = "ab"[r.Intn(2)]
		}
		for i := range tt {
			tt[i] = "ab"[r.Intn(2)]
		}
		want := Recursive(string(s), string(tt))
		if Tab(string(s), string(tt)) != want || Optimal(string(s), string(tt)) != want {
			t.Fatalf("%q %q", s, tt)
		}
	}
}
