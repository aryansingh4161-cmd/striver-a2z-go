package triangle

import (
	"math/rand"
	"testing"
)

func TestTriangle(t *testing.T) {
	fs := []func([][]int) int{Recursive, Memo, Tab, Optimal}
	cases := []struct {
		tri  [][]int
		want int
	}{
		{[][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}, 11},
		{[][]int{{-10}}, -10},
		{[][]int{{1}, {2, 3}}, 3},
	}
	for _, c := range cases {
		for i, f := range fs {
			if got := f(c.tri); got != c.want {
				t.Errorf("f%d = %d want %d", i, got, c.want)
			}
		}
	}
	r := rand.New(rand.NewSource(9))
	for it := 0; it < 200; it++ {
		n := 1 + r.Intn(8)
		tri := make([][]int, n)
		for i := range tri {
			tri[i] = make([]int, i+1)
			for j := range tri[i] {
				tri[i][j] = r.Intn(21) - 10
			}
		}
		want := Recursive(tri)
		for i, f := range fs[1:] {
			if f(tri) != want {
				t.Fatalf("f%d on %v", i+1, tri)
			}
		}
	}
}
