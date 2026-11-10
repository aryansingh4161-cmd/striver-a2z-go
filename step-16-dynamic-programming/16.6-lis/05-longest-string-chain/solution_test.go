package longeststringchain

import (
	"math/rand"
	"testing"
)

func TestChain(t *testing.T) {
	cases := []struct {
		w    []string
		want int
	}{
		{[]string{"a", "b", "ba", "bca", "bda", "bdca"}, 4},
		{[]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}, 5},
		{[]string{"abcd", "dbqca"}, 1},
	}
	for _, c := range cases {
		if LongestChain(c.w) != c.want || LongestChainMap(c.w) != c.want {
			t.Errorf("%v", c.w)
		}
	}
	r := rand.New(rand.NewSource(5))
	for it := 0; it < 200; it++ {
		seen := map[string]bool{}
		var w []string
		for k := 1 + r.Intn(10); k > 0; k-- {
			b := make([]byte, 1+r.Intn(4))
			for i := range b {
				b[i] = "ab"[r.Intn(2)]
			}
			if !seen[string(b)] {
				seen[string(b)] = true
				w = append(w, string(b))
			}
		}
		if LongestChain(w) != LongestChainMap(w) {
			t.Fatalf("%v", w)
		}
	}
}
