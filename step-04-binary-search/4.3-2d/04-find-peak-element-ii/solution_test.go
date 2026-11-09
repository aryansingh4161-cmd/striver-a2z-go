package findpeakelementii

import (
	"math/rand"
	"testing"
)

func isPeak(m [][]int, r, c int) bool {
	v := m[r][c]
	for _, d := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		i, j := r+d[0], c+d[1]
		if i >= 0 && i < len(m) && j >= 0 && j < len(m[0]) && m[i][j] >= v {
			return false
		}
	}
	return true
}

func TestFindPeak(t *testing.T) {
	cases := [][][]int{
		{{1, 4}, {3, 2}},
		{{10, 20, 15}, {21, 30, 14}, {7, 16, 32}},
		{{5}},
	}
	// random grids with distinct values
	r := rand.New(rand.NewSource(2))
	for k := 0; k < 50; k++ {
		n, m := r.Intn(5)+1, r.Intn(5)+1
		perm := r.Perm(n * m)
		g := make([][]int, n)
		for i := range g {
			g[i] = perm[i*m : (i+1)*m]
		}
		cases = append(cases, g)
	}
	for _, g := range cases {
		p := FindPeak(g)
		if p[0] < 0 || !isPeak(g, p[0], p[1]) {
			t.Fatalf("FindPeak(%v) = %v not a peak", g, p)
		}
	}
}
