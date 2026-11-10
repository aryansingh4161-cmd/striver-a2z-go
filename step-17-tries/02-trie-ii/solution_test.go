package trieii

import (
	"math/rand"
	"strings"
	"testing"
)

func TestTrieII(t *testing.T) {
	tr := New()
	for _, w := range []string{"apple", "apple", "apps", "apxl"} {
		tr.Insert(w)
	}
	if tr.CountWordsEqualTo("apple") != 2 || tr.CountWordsStartingWith("app") != 3 || tr.CountWordsStartingWith("ap") != 4 {
		t.Error("counts")
	}
	tr.Erase("apple")
	if tr.CountWordsEqualTo("apple") != 1 || tr.CountWordsStartingWith("app") != 2 || tr.CountWordsEqualTo("zz") != 0 {
		t.Error("after erase")
	}
}

func TestAgainstMap(t *testing.T) {
	r := rand.New(rand.NewSource(2))
	tr, cnt := New(), map[string]int{}
	for it := 0; it < 3000; it++ {
		b := make([]byte, 1+r.Intn(3))
		for i := range b {
			b[i] = "ab"[r.Intn(2)]
		}
		w := string(b)
		switch r.Intn(4) {
		case 0:
			tr.Insert(w)
			cnt[w]++
		case 1:
			if cnt[w] > 0 {
				tr.Erase(w)
				cnt[w]--
			}
		case 2:
			if tr.CountWordsEqualTo(w) != cnt[w] {
				t.Fatalf("equal %q", w)
			}
		default:
			want := 0
			for s, c := range cnt {
				if strings.HasPrefix(s, w) {
					want += c
				}
			}
			if tr.CountWordsStartingWith(w) != want {
				t.Fatalf("prefix %q", w)
			}
		}
	}
}
