package longestcompletestring

import (
	"math/rand"
	"testing"
)

func TestCompleteString(t *testing.T) {
	cases := []struct {
		w    []string
		want string
	}{
		{[]string{"n", "ni", "nin", "ninj", "ninja", "ninga"}, "ninja"},
		{[]string{"ab", "bc"}, "None"},
		{[]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}, "apple"},
		{[]string{"b", "a"}, "a"},
	}
	for _, c := range cases {
		if got := CompleteString(c.w); got != c.want {
			t.Errorf("%v got %q", c.w, got)
		}
	}
	r := rand.New(rand.NewSource(3))
	for it := 0; it < 300; it++ {
		w := make([]string, 1+r.Intn(10))
		for i := range w {
			b := make([]byte, 1+r.Intn(4))
			for j := range b {
				b[j] = "ab"[r.Intn(2)]
			}
			w[i] = string(b)
		}
		if CompleteString(w) != Brute(w) {
			t.Fatalf("%v", w)
		}
	}
}
