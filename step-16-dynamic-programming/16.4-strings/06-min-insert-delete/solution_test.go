package mininsertdelete

import "testing"

func TestMinOperations(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{{"abcd", "anc", 3}, {"sea", "eat", 2}, {"leetcode", "etco", 4}, {"", "ab", 2}, {"x", "x", 0}}
	for _, c := range cases {
		if got := MinOperations(c.a, c.b); got != c.want {
			t.Errorf("%q→%q got %d want %d", c.a, c.b, got, c.want)
		}
	}
}
