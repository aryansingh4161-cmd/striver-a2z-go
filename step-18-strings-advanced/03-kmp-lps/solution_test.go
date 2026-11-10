package kmplps

import (
	"math/rand"
	"slices"
	"testing"
)

func TestKMP(t *testing.T) {
	if got := LPS("aabaaa"); !slices.Equal(got, []int{0, 1, 0, 1, 2, 2}) {
		t.Errorf("lps %v", got)
	}
	if got := LPS("abcabcd"); !slices.Equal(got, []int{0, 0, 0, 1, 2, 3, 0}) {
		t.Errorf("lps %v", got)
	}
	if got := Search("aaaa", "aa"); !slices.Equal(got, []int{0, 1, 2}) {
		t.Errorf("overlaps %v", got)
	}
	r := rand.New(rand.NewSource(3))
	for it := 0; it < 500; it++ {
		s, p := make([]byte, r.Intn(25)), make([]byte, 1+r.Intn(5))
		for i := range s {
			s[i] = "ab"[r.Intn(2)]
		}
		for i := range p {
			p[i] = "ab"[r.Intn(2)]
		}
		if !slices.Equal(LPS(string(p)), LPSBrute(string(p))) {
			t.Fatalf("lps %q", p)
		}
		var want []int
		for i := 0; i+len(p) <= len(s); i++ {
			if string(s[i:i+len(p)]) == string(p) {
				want = append(want, i)
			}
		}
		if got := Search(string(s), string(p)); !slices.Equal(got, want) && !(len(got) == 0 && len(want) == 0) {
			t.Fatalf("%q in %q: %v vs %v", p, s, got, want)
		}
	}
}
