package matrixmedian

import (
	"math/rand"
	"slices"
	"testing"
)

func TestMedian(t *testing.T) {
	if got := Median([][]int{{1, 4, 9}, {2, 5, 6}, {3, 7, 8}}); got != 5 {
		t.Errorf("got %d", got)
	}
	if got := Median([][]int{{1, 3, 5}, {2, 6, 9}, {3, 6, 9}}); got != 5 {
		t.Errorf("got %d", got)
	}
	r := rand.New(rand.NewSource(4))
	for k := 0; k < 100; k++ {
		n, m := 2*r.Intn(3)+1, 2*r.Intn(3)+1 // odd * odd = odd
		g := make([][]int, n)
		var all []int
		for i := range g {
			g[i] = make([]int, m)
			for j := range g[i] {
				g[i][j] = r.Intn(30)
			}
			slices.Sort(g[i])
			all = append(all, g[i]...)
		}
		slices.Sort(all)
		if got, want := Median(g), all[len(all)/2]; got != want {
			t.Fatalf("Median(%v) = %d want %d", g, got, want)
		}
	}
}
