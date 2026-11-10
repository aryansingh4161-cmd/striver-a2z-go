package longestcommonsubstring

import (
	"math/rand"
	"testing"
)

func TestSubstring(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{{"abcjklp", "acjkp", 3}, {"abc", "xyz", 0}, {"aaaa", "aa", 2}, {"", "a", 0}}
	for _, c := range cases {
		for i, f := range []func(string, string) int{Brute, Tab, Optimal} {
			if got := f(c.a, c.b); got != c.want {
				t.Errorf("f%d(%q,%q) = %d", i, c.a, c.b, got)
			}
		}
	}
	r := rand.New(rand.NewSource(3))
	for it := 0; it < 300; it++ {
		a, b := make([]byte, r.Intn(10)), make([]byte, r.Intn(10))
		for i := range a {
			a[i] = "ab"[r.Intn(2)]
		}
		for i := range b {
			b[i] = "ab"[r.Intn(2)]
		}
		want := Brute(string(a), string(b))
		if Tab(string(a), string(b)) != want || Optimal(string(a), string(b)) != want {
			t.Fatalf("%q %q", a, b)
		}
	}
}
