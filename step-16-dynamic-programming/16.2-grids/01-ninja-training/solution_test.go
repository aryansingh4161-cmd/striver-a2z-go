package ninjatraining

import (
	"math/rand"
	"testing"
)

func TestNinja(t *testing.T) {
	fs := []func([][3]int) int{Recursive, Memo, Tab, Optimal}
	cases := []struct {
		p    [][3]int
		want int
	}{
		{[][3]int{{1, 2, 5}, {3, 1, 1}, {3, 3, 3}}, 11},
		{[][3]int{{10, 50, 1}, {5, 100, 11}}, 110},
		{[][3]int{{10, 40, 70}, {20, 50, 80}, {30, 60, 90}}, 210},
		{[][3]int{{4, 9, 2}}, 9},
	}
	for _, c := range cases {
		for i, f := range fs {
			if got := f(c.p); got != c.want {
				t.Errorf("f%d = %d want %d", i, got, c.want)
			}
		}
	}
	r := rand.New(rand.NewSource(6))
	for it := 0; it < 200; it++ {
		p := make([][3]int, 1+r.Intn(8))
		for i := range p {
			p[i] = [3]int{r.Intn(50), r.Intn(50), r.Intn(50)}
		}
		want := Recursive(p)
		for i, f := range fs[1:] {
			if f(p) != want {
				t.Fatalf("f%d mismatch on %v", i+1, p)
			}
		}
	}
}
