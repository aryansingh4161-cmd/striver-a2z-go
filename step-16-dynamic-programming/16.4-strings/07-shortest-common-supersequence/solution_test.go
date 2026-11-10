package shortestcommonsupersequence

import (
	"math/rand"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-16-dynamic-programming/16.4-strings/01-lcs"
)

func isSubseq(s, t string) bool {
	i := 0
	for j := 0; j < len(t) && i < len(s); j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == len(s)
}

func TestSCS(t *testing.T) {
	if got := SCS("brute", "groot"); got != "bgruoote" {
		t.Errorf("got %q", got)
	}
	if got := SCS("abac", "cab"); len(got) != 5 || !isSubseq("abac", got) || !isSubseq("cab", got) {
		t.Errorf("got %q", got)
	}
	r := rand.New(rand.NewSource(7))
	for it := 0; it < 300; it++ {
		a, b := make([]byte, r.Intn(8)), make([]byte, r.Intn(8))
		for i := range a {
			a[i] = "abc"[r.Intn(3)]
		}
		for i := range b {
			b[i] = "abc"[r.Intn(3)]
		}
		got := SCS(string(a), string(b))
		want := len(a) + len(b) - lcs.Optimal(string(a), string(b))
		if len(got) != want || !isSubseq(string(a), got) || !isSubseq(string(b), got) {
			t.Fatalf("%q %q → %q", a, b, got)
		}
	}
}
