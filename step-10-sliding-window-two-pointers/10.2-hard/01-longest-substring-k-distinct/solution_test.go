package longestsubstringkdistinct

import (
	"math/rand"
	"testing"
)

func TestLongest(t *testing.T) {
	if Longest("aaabbccd", 2) != 5 || Longest("eceba", 2) != 3 || Longest("aa", 1) != 2 || Longest("abc", 0) != 0 {
		t.Error("examples")
	}
	r := rand.New(rand.NewSource(6))
	for it := 0; it < 300; it++ {
		b := make([]byte, r.Intn(20))
		for i := range b {
			b[i] = byte('a' + r.Intn(5))
		}
		k := r.Intn(5)
		if Longest(string(b), k) != LongestBrute(string(b), k) {
			t.Fatalf("mismatch %q k=%d", b, k)
		}
	}
}
