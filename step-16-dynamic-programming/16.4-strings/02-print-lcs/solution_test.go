package printlcs

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

func TestPrintLCS(t *testing.T) {
	if got := PrintLCS("abcde", "bdgek"); got != "bde" {
		t.Errorf("got %q", got)
	}
	r := rand.New(rand.NewSource(2))
	for it := 0; it < 300; it++ {
		a, b := make([]byte, r.Intn(9)), make([]byte, r.Intn(9))
		for i := range a {
			a[i] = "abc"[r.Intn(3)]
		}
		for i := range b {
			b[i] = "abc"[r.Intn(3)]
		}
		got := PrintLCS(string(a), string(b))
		if len(got) != lcs.Optimal(string(a), string(b)) || !isSubseq(got, string(a)) || !isSubseq(got, string(b)) {
			t.Fatalf("%q %q → %q", a, b, got)
		}
	}
}
