package countsubstringskdistinct

import (
	"math/rand"
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		s       string
		k, want int
	}{{"pqpqs", 2, 7}, {"aabab", 3, 0}, {"abc", 1, 3}, {"aa", 1, 3}}
	for _, tc := range tests {
		if got := Count(tc.s, tc.k); got != tc.want {
			t.Errorf("Count(%q,%d) = %d", tc.s, tc.k, got)
		}
	}
	r := rand.New(rand.NewSource(1))
	for it := 0; it < 200; it++ {
		b := make([]byte, r.Intn(12))
		for i := range b {
			b[i] = byte('a' + r.Intn(4))
		}
		k := r.Intn(5)
		if Count(string(b), k) != CountBrute(string(b), k) {
			t.Fatalf("mismatch %q k=%d", b, k)
		}
	}
}
