package trie

import (
	"math/rand"
	"strings"
	"testing"
)

func TestTrie(t *testing.T) {
	tr := New()
	tr.Insert("apple")
	if !tr.Search("apple") || tr.Search("app") || !tr.StartsWith("app") || tr.StartsWith("b") {
		t.Error("after apple")
	}
	tr.Insert("app")
	if !tr.Search("app") {
		t.Error("app inserted")
	}
	if !tr.StartsWith("") {
		t.Error("empty prefix matches everything")
	}
}

func TestAgainstMap(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	word := func() string {
		b := make([]byte, 1+r.Intn(4))
		for i := range b {
			b[i] = "abc"[r.Intn(3)]
		}
		return string(b)
	}
	tr, set := New(), map[string]bool{}
	for it := 0; it < 2000; it++ {
		w := word()
		switch r.Intn(3) {
		case 0:
			tr.Insert(w)
			set[w] = true
		case 1:
			if tr.Search(w) != set[w] {
				t.Fatalf("search %q", w)
			}
		default:
			want := false
			for s := range set {
				if strings.HasPrefix(s, w) {
					want = true
				}
			}
			if tr.StartsWith(w) != want {
				t.Fatalf("startsWith %q", w)
			}
		}
	}
}
