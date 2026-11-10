package minimumwindowsubsequence

import (
	"math/rand"
	"testing"
)

func TestMinWindow(t *testing.T) {
	if got := MinWindow("abcdebdde", "bde"); got != "bcde" {
		t.Errorf("got %q", got)
	}
	if got := MinWindow("jmeqksfrsdcmsiwvaovztaqenprpvnbstl", "u"); got != "" {
		t.Errorf("got %q", got)
	}
	r := rand.New(rand.NewSource(8))
	for it := 0; it < 500; it++ {
		a := make([]byte, r.Intn(15))
		b := make([]byte, r.Intn(4)+1)
		for i := range a {
			a[i] = byte('a' + r.Intn(3))
		}
		for i := range b {
			b[i] = byte('a' + r.Intn(3))
		}
		if MinWindow(string(a), string(b)) != MinWindowBrute(string(a), string(b)) {
			t.Fatalf("mismatch %q %q: %q vs %q", a, b, MinWindow(string(a), string(b)), MinWindowBrute(string(a), string(b)))
		}
	}
}
