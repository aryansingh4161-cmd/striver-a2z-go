package zfunction

import (
	"math/rand"
	"slices"
	"strings"
	"testing"
)

func TestZ(t *testing.T) {
	if got := Z("aabxaab"); !slices.Equal(got, []int{0, 1, 0, 0, 3, 1, 0}) {
		t.Errorf("got %v", got)
	}
	if got := Search("abababab", "aba"); !slices.Equal(got, []int{0, 2, 4}) {
		t.Errorf("got %v", got)
	}
	r := rand.New(rand.NewSource(2))
	for it := 0; it < 500; it++ {
		b := make([]byte, r.Intn(25))
		for i := range b {
			b[i] = "ab"[r.Intn(2)]
		}
		if !slices.Equal(Z(string(b)), ZBrute(string(b))) {
			t.Fatalf("%q", b)
		}
		p := string(b[:min(len(b), 1+r.Intn(3))])
		if p == "" {
			continue
		}
		var want []int
		for i := 0; i+len(p) <= len(b); i++ {
			if strings.HasPrefix(string(b[i:]), p) {
				want = append(want, i)
			}
		}
		if got := Search(string(b), p); !slices.Equal(got, want) && !(len(got) == 0 && len(want) == 0) {
			t.Fatalf("search %q in %q: %v vs %v", p, b, got, want)
		}
	}
}
