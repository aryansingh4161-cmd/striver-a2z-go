package longestsubstringnorepeat

import (
	"math/rand"
	"testing"
)

func TestLongest(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{{"abcabcbb", 3}, {"bbbbb", 1}, {"pwwkew", 3}, {"", 0}, {"abba", 2}, {" ", 1}}
	for _, tc := range tests {
		if got := Longest(tc.s); got != tc.want {
			t.Errorf("Longest(%q) = %d", tc.s, got)
		}
	}
	r := rand.New(rand.NewSource(1))
	for k := 0; k < 300; k++ {
		b := make([]byte, r.Intn(20))
		for i := range b {
			b[i] = byte('a' + r.Intn(5))
		}
		if Longest(string(b)) != LongestBrute(string(b)) {
			t.Fatalf("mismatch %q", b)
		}
	}
}
