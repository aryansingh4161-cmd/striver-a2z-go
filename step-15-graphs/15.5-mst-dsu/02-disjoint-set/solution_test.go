package disjointset

import (
	"math/rand"
	"testing"
)

func TestDSU(t *testing.T) {
	d := New(8)
	for _, p := range [][2]int{{1, 2}, {2, 3}, {4, 5}, {6, 7}, {5, 6}} {
		d.Union(p[0], p[1])
	}
	if d.Same(3, 7) {
		t.Error("3 and 7 not yet joined")
	}
	d.Union(3, 7)
	if !d.Same(3, 7) || d.Size(1) != 7 || d.Groups != 2 {
		t.Errorf("same=%v size=%d groups=%d", d.Same(3, 7), d.Size(1), d.Groups)
	}
	if d.Union(1, 7) {
		t.Error("already joined → false")
	}
}

// Compare both DSUs with a naive "label every item" version.
func TestRandom(t *testing.T) {
	r := rand.New(rand.NewSource(13))
	for it := 0; it < 100; it++ {
		n := 1 + r.Intn(20)
		d, rd := New(n), NewRank(n)
		label := make([]int, n)
		for i := range label {
			label[i] = i
		}
		for op := 0; op < 40; op++ {
			a, b := r.Intn(n), r.Intn(n)
			if r.Intn(2) == 0 {
				d.Union(a, b)
				rd.Union(a, b)
				la, lb := label[a], label[b]
				for i := range label {
					if label[i] == lb {
						label[i] = la
					}
				}
			} else {
				want := label[a] == label[b]
				if d.Same(a, b) != want || (rd.Find(a) == rd.Find(b)) != want {
					t.Fatal("mismatch")
				}
			}
		}
	}
}
