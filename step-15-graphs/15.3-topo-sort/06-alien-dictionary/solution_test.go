package aliendictionary

import (
	"math/rand"
	"sort"
	"strings"
	"testing"
)

// sortedUnder reports whether words are sorted under the given alphabet.
func sortedUnder(words []string, alpha string) bool {
	rank := map[byte]int{}
	for i := 0; i < len(alpha); i++ {
		rank[alpha[i]] = i
	}
	less := func(a, b string) bool {
		for i := 0; i < len(a) && i < len(b); i++ {
			if a[i] != b[i] {
				return rank[a[i]] < rank[b[i]]
			}
		}
		return len(a) <= len(b)
	}
	for i := 0; i+1 < len(words); i++ {
		if !less(words[i], words[i+1]) {
			return false
		}
	}
	return true
}

func TestFindOrder(t *testing.T) {
	if got := FindOrder([]string{"baa", "abcd", "abca", "cab", "cad"}, 4); got != "bdac" {
		t.Errorf("got %q", got)
	}
	if got := FindOrder([]string{"caa", "aaa", "aab"}, 3); !sortedUnder([]string{"caa", "aaa", "aab"}, got) || len(got) != 3 {
		t.Errorf("got %q", got)
	}
	if FindOrder([]string{"abc", "ab"}, 3) != "" {
		t.Error("prefix case must fail")
	}
	if FindOrder([]string{"ab", "ba", "ab"}, 2) != "" {
		t.Error("cycle must fail")
	}
}

// Build words sorted under a secret alphabet; the answer must sort them too.
func TestRandom(t *testing.T) {
	r := rand.New(rand.NewSource(6))
	for it := 0; it < 200; it++ {
		k := 1 + r.Intn(5)
		secret := []byte("abcde"[:k])
		r.Shuffle(k, func(i, j int) { secret[i], secret[j] = secret[j], secret[i] })
		rank := map[byte]int{}
		for i, c := range secret {
			rank[c] = i
		}
		words := make([]string, 2+r.Intn(6))
		for i := range words {
			var sb strings.Builder
			for l := 1 + r.Intn(4); l > 0; l-- {
				sb.WriteByte(byte('a' + r.Intn(k)))
			}
			words[i] = sb.String()
		}
		sort.Slice(words, func(i, j int) bool {
			a, b := words[i], words[j]
			for x := 0; x < len(a) && x < len(b); x++ {
				if a[x] != b[x] {
					return rank[a[x]] < rank[b[x]]
				}
			}
			return len(a) < len(b)
		})
		got := FindOrder(words, k)
		if len(got) != k || !sortedUnder(words, got) {
			t.Fatalf("words %v secret %s got %q", words, secret, got)
		}
	}
}
