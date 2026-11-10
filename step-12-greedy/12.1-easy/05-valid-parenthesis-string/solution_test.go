package validparenthesisstring

import (
	"math/rand"
	"testing"
)

func TestCheckValidString(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{{"()", true}, {"(*)", true}, {"(*))", true}, {"(()", false}, {"*(", false}, {"", true}}
	for _, tc := range tests {
		if got := CheckValidString(tc.s); got != tc.want {
			t.Errorf("CheckValidString(%q) = %v", tc.s, got)
		}
	}
	r := rand.New(rand.NewSource(1))
	for it := 0; it < 500; it++ {
		b := make([]byte, r.Intn(12))
		for i := range b {
			b[i] = "()*"[r.Intn(3)]
		}
		if CheckValidString(string(b)) != CheckValidStringBrute(string(b)) {
			t.Fatalf("mismatch %q", b)
		}
	}
}
