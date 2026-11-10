package substringsallthreechars

import (
	"math/rand"
	"testing"
)

func TestNumberOfSubstrings(t *testing.T) {
	if NumberOfSubstrings("abcabc") != 10 || NumberOfSubstrings("aaacb") != 3 || NumberOfSubstrings("abc") != 1 {
		t.Error("examples")
	}
	r := rand.New(rand.NewSource(5))
	for it := 0; it < 300; it++ {
		b := make([]byte, r.Intn(20))
		for i := range b {
			b[i] = byte('a' + r.Intn(3))
		}
		if NumberOfSubstrings(string(b)) != NumberOfSubstringsBrute(string(b)) {
			t.Fatalf("mismatch %q", b)
		}
	}
}
