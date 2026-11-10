package rabinkarp

import (
	"math/rand"
	"slices"
	"testing"
)

func TestSearch(t *testing.T) {
	if got := Search("abababab", "aba"); !slices.Equal(got, []int{0, 2, 4}) {
		t.Errorf("got %v", got)
	}
	if got := Search("abc", "abcd"); len(got) != 0 {
		t.Errorf("pattern longer than text: %v", got)
	}
	r := rand.New(rand.NewSource(1))
	for it := 0; it < 500; it++ {
		s, p := make([]byte, r.Intn(30)), make([]byte, 1+r.Intn(4))
		for i := range s {
			s[i] = "ab"[r.Intn(2)]
		}
		for i := range p {
			p[i] = "ab"[r.Intn(2)]
		}
		if got, want := Search(string(s), string(p)), Brute(string(s), string(p)); !slices.Equal(got, want) {
			t.Fatalf("%q %q: %v vs %v", s, p, got, want)
		}
	}
}
