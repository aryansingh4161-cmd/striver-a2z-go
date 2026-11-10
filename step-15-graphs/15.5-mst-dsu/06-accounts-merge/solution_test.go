package accountsmerge

import (
	"sort"
	"strings"
	"testing"
)

func canon(a [][]string) []string {
	s := make([]string, len(a))
	for i, x := range a {
		s[i] = strings.Join(x, ",")
	}
	sort.Strings(s)
	return s
}

func TestMerge(t *testing.T) {
	in := [][]string{
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"Mary", "mary@mail.com"},
		{"John", "johnnybravo@mail.com"},
	}
	want := canon([][]string{
		{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
		{"Mary", "mary@mail.com"},
		{"John", "johnnybravo@mail.com"},
	})
	got := canon(Merge(in))
	if strings.Join(got, "|") != strings.Join(want, "|") {
		t.Errorf("got %v", got)
	}
	// chain: 0-1 share a, 1-2 share b → all one person
	chain := [][]string{{"A", "a", "x"}, {"A", "a", "b"}, {"A", "b", "c"}}
	if g := Merge(chain); len(g) != 1 || len(g[0]) != 5 {
		t.Errorf("chain got %v", g)
	}
}
