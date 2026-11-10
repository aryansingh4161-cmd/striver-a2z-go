package editdistance

import (
	"math/rand"
	"testing"
)

func TestEditDistance(t *testing.T) {
	fs := []func(string, string) int{Recursive, Memo, Optimal}
	cases := []struct {
		a, b string
		want int
	}{{"horse", "ros", 3}, {"intention", "execution", 5}, {"", "abc", 3}, {"abc", "", 3}, {"same", "same", 0}}
	for _, c := range cases {
		for i, f := range fs {
			if got := f(c.a, c.b); got != c.want {
				t.Errorf("f%d(%q,%q) = %d want %d", i, c.a, c.b, got, c.want)
			}
		}
	}
	r := rand.New(rand.NewSource(9))
	for it := 0; it < 200; it++ {
		a, b := make([]byte, r.Intn(6)), make([]byte, r.Intn(6))
		for i := range a {
			a[i] = "abc"[r.Intn(3)]
		}
		for i := range b {
			b[i] = "abc"[r.Intn(3)]
		}
		want := Recursive(string(a), string(b))
		if Memo(string(a), string(b)) != want || Optimal(string(a), string(b)) != want {
			t.Fatalf("%q %q", a, b)
		}
	}
}
