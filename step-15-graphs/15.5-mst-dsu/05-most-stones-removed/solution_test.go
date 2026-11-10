package moststonesremoved

import (
	"math/rand"
	"testing"
)

// brute: union stones pairwise, O(n²).
func brute(st [][2]int) int {
	n := len(st)
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if p[x] != x {
			p[x] = find(p[x])
		}
		return p[x]
	}
	groups := n
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if st[i][0] == st[j][0] || st[i][1] == st[j][1] {
				if a, b := find(i), find(j); a != b {
					p[a] = b
					groups--
				}
			}
		}
	}
	return n - groups
}

func TestRemoveStones(t *testing.T) {
	if got := RemoveStones([][2]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}); got != 5 {
		t.Errorf("got %d", got)
	}
	if got := RemoveStones([][2]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}); got != 3 {
		t.Errorf("got %d", got)
	}
	if got := RemoveStones([][2]int{{0, 0}}); got != 0 {
		t.Errorf("got %d", got)
	}
	r := rand.New(rand.NewSource(15))
	for it := 0; it < 300; it++ {
		seen := map[[2]int]bool{}
		var st [][2]int
		for k := 1 + r.Intn(10); k > 0; k-- {
			s := [2]int{r.Intn(6), r.Intn(6)}
			if !seen[s] {
				seen[s] = true
				st = append(st, s)
			}
		}
		if a, b := RemoveStones(st), brute(st); a != b {
			t.Fatalf("%v: %d vs %d", st, a, b)
		}
	}
}
